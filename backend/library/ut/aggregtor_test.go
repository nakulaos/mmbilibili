package ut

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestAggregator_Basic(t *testing.T) {
	batchProcess := func(items []int) error {
		t.Logf("handler %d items", len(items))
		return nil
	}

	aggregator := NewAggregator(batchProcess)

	aggregator.Start()

	for i := 0; i < 1000; i++ {
		aggregator.TryEnqueue(i)
	}

	aggregator.SafeStop()
}

func TestAggregator_Complex(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)

	batchProcess := func(items []int) error {
		defer wg.Add(-len(items))
		time.Sleep(20 * time.Millisecond)
		if len(items) != 4 {
			return errors.New("len(items) != 4")
		}
		return nil
	}

	errorHandler := func(err error, items []int, batchProcessFunc BatchProcessFunc[int], aggregator *Aggregator[int]) {
		if err == nil {
			t.FailNow()
		}
		t.Logf("Receive error, item size is %d", len(items))
	}

	aggregator := NewAggregator(batchProcess, func(option AggregatorOption[int]) AggregatorOption[int] {
		option.BatchSize = 4
		option.Workers = 1
		option.ErrorHandler = errorHandler
		option.Logger = NewConsoleLogger()
		return option
	})

	aggregator.Start()

	for i := 0; i < 100; i++ {
		for !aggregator.TryEnqueue(i) {
			time.Sleep(10 * time.Millisecond)
		}
	}

	aggregator.SafeStop()
	wg.Wait()
}

func TestAggregator_LingerTimeOut(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)

	batchProcess := func(items []int) error {
		defer wg.Add(-len(items))
		if len(items) != 4 {
			t.Log("linger time out")
		}
		return nil
	}

	aggregator := NewAggregator(batchProcess, func(option AggregatorOption[int]) AggregatorOption[int] {
		option.BatchSize = 4
		option.Workers = 1
		option.LingerTime = 100 * time.Millisecond
		option.Logger = NewConsoleLogger()
		return option
	})

	aggregator.Start()

	for i := 0; i < 100; i++ {
		aggregator.TryEnqueue(i)
		time.Sleep(100 * time.Millisecond)
	}

	aggregator.SafeStop()
	wg.Wait()
}
