package ut

type Promise struct {
	done chan struct{} // done 指示异步操作是否完成
	res  interface{}   // res 保存异步操作的结果
	err  error         // err 保存异步操作的错误
}

type WorkFunc func() (interface{}, error)

type SuccessHandler func(interface{}) (interface{}, error)

type ErrorHandler func(error) interface{}

func NewPromise(workFunc WorkFunc) *Promise {
	promise := Promise{done: make(chan struct{})}
	go func() {
		// 异步执行 workFunc
		// 执行完之后关闭 done 通道
		defer close(promise.done)
		promise.res, promise.err = workFunc()
	}()
	return &promise
}

// Done 等待异步操作完成，返回结果和错误
func (p *Promise) Done() (interface{}, error) {
	<-p.done
	return p.res, p.err
}

// Then 设置回调函数，当异步操作完成时，调用回调函数
func (p *Promise) Then(successHandler SuccessHandler, errorHandler ErrorHandler) *Promise {
	newPromise := &Promise{done: make(chan struct{})}
	go func() {
		res, err := p.Done()
		defer close(newPromise.done)
		if err != nil {
			if errorHandler != nil {
				newPromise.res = errorHandler(err)
			} else {
				newPromise.err = err
			}
		} else {
			if successHandler != nil {
				newPromise.res, newPromise.err = successHandler(res)
			} else {
				newPromise.res = res
			}
		}
	}()

	return newPromise
}

func (p *Promise) ThenSuccess(successHandler SuccessHandler) *Promise {
	return p.Then(successHandler, nil)
}

func (p *Promise) ThenError(errorHandler ErrorHandler) *Promise {
	return p.Then(nil, errorHandler)
}
