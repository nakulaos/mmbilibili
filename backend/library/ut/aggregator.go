package ut

import (
	"runtime"
	"sync"
	"time"
)

type Aggregator[T any] struct {
	option         AggregatorOption[T]
	wg             *sync.WaitGroup
	quit           chan struct{}
	eventQueue     chan T
	batchProcessor BatchProcessFunc[T]
}

type AggregatorOption[T any] struct {
	BatchSize         int                 // 每批次达到多少才处理
	Workers           int                 // 处理批次的并发数
	ChannelBufferSize int                 // 事件队列的缓冲大小
	LingerTime        time.Duration       // 事件队列的等待时间
	ErrorHandler      ErrorHandlerFunc[T] // 错误处理函数
	Logger            Logger              // 日志
}

type BatchProcessFunc[T any] func([]T) error

type SetAggregatorOptionFunc[T any] func(option AggregatorOption[T]) AggregatorOption[T]

type ErrorHandlerFunc[T any] func(err error, items []T, batchProcessFunc BatchProcessFunc[T], aggregator *Aggregator[T])

func NewAggregator[T any](batchProcessor BatchProcessFunc[T], optionFuncs ...SetAggregatorOptionFunc[T]) *Aggregator[T] {
	option := AggregatorOption[T]{
		BatchSize:  8,
		Workers:    runtime.NumCPU(),
		LingerTime: 1 * time.Minute,
	}

	for _, optionFunc := range optionFuncs {
		option = optionFunc(option)
	}

	if option.ChannelBufferSize <= option.Workers {
		option.ChannelBufferSize = option.Workers
	}

	return &Aggregator[T]{
		eventQueue:     make(chan T, option.ChannelBufferSize),
		option:         option,
		quit:           make(chan struct{}),
		wg:             new(sync.WaitGroup),
		batchProcessor: batchProcessor,
	}
}

func (agt *Aggregator[T]) TryEnqueue(item T) bool {
	select {
	case agt.eventQueue <- item:
		return true
	default:
		if agt.option.Logger != nil {
			agt.option.Logger.Warnf("Aggregator: Event queue is full and try reschedule")
		}

		runtime.Gosched()

		select {
		case agt.eventQueue <- item:
			return true
		default:
			if agt.option.Logger != nil {
				agt.option.Logger.Warnf("Aggregator: Event queue is still full and %+v is skipped.", item)
			}
			return false
		}
	}
}

func (agt *Aggregator[T]) Enqueue(item T) {
	agt.eventQueue <- item
}

func (agt *Aggregator[T]) Start() {
	for i := 0; i < agt.option.Workers; i++ {
		index := i
		go agt.work(index)
	}
}

func (agt *Aggregator[T]) Stop() {
	close(agt.quit)
	agt.wg.Wait()
}

func (agt *Aggregator[T]) SafeStop() {
	if len(agt.eventQueue) == 0 {
		close(agt.quit)
	} else {
		ticker := time.NewTicker(50 * time.Millisecond)
		for range ticker.C {
			if len(agt.eventQueue) == 0 {
				close(agt.quit)
				break
			}
		}
		ticker.Stop()
	}
	agt.wg.Wait()
}

func (agt *Aggregator[T]) work(index int) {
	defer func() {
		if r := recover(); r != nil {
			if agt.option.Logger != nil {
				agt.option.Logger.Errorf("Aggregator: recover worker as bad thing happens %+v", r)
			}

			agt.work(index)
		}
	}()

	agt.wg.Add(1)
	defer agt.wg.Done()

	batch := make([]T, 0, agt.option.BatchSize)
	lingerTimer := time.NewTimer(0)
	if !lingerTimer.Stop() {
		<-lingerTimer.C
	}
	defer lingerTimer.Stop()

loop:
	for {
		select {
		case req := <-agt.eventQueue:
			batch = append(batch, req)

			batchSize := len(batch)
			if batchSize < agt.option.BatchSize {
				if batchSize == 1 {
					lingerTimer.Reset(agt.option.LingerTime)
				}
				break
			}

			agt.batchProcess(batch)

			if !lingerTimer.Stop() {
				<-lingerTimer.C
			}
			batch = make([]T, 0, agt.option.BatchSize)
		case <-lingerTimer.C:
			if len(batch) == 0 {
				break
			}

			agt.batchProcess(batch)
			batch = make([]T, 0, agt.option.BatchSize)
		case <-agt.quit:
			if len(batch) != 0 {
				agt.batchProcess(batch)
			}

			break loop
		}
	}
}

func (agt *Aggregator[T]) batchProcess(items []T) {
	agt.wg.Add(1)
	defer agt.wg.Done()
	if err := agt.batchProcessor(items); err != nil {
		if agt.option.Logger != nil {
			agt.option.Logger.Errorf("Aggregator: error happens")
		}

		if agt.option.ErrorHandler != nil {
			go agt.option.ErrorHandler(err, items, agt.batchProcessor, agt)
		} else if agt.option.Logger != nil {
			agt.option.Logger.Errorf("Aggregator: error happens in batchProcess and is skipped")
		}
	} else if agt.option.Logger != nil {
		agt.option.Logger.Infof("Aggregator: %d items have been sent.", len(items))
	}
}

// 添加配置函数

func SetAggregatorBatchSize[T any](size int) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.BatchSize = size
		return option
	}
}

func SetAggregatorWorkers[T any](workers int) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.Workers = workers
		return option
	}
}

func SetAggregatorChannelBufferSize[T any](size int) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.ChannelBufferSize = size
		return option
	}
}

func SetAggregatorLingerTime[T any](lingerTime time.Duration) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.LingerTime = lingerTime
		return option
	}
}

func SetAggregatorErrorHandler[T any](errorHandler ErrorHandlerFunc[T]) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.ErrorHandler = errorHandler
		return option
	}
}

func SetAggregatorLogger[T any](logger Logger) SetAggregatorOptionFunc[T] {
	return func(option AggregatorOption[T]) AggregatorOption[T] {
		option.Logger = logger
		return option
	}
}
