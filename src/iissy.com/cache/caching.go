package cache

// Func 是一个函数类型
// 从数据源取数据的函数就必须符合这个函数类型的签名
type Func func(key int) ([]string, error)

// 这个结构是缓存的数据，value存放用户需要的缓存数据
type result struct {
	value []string
	err   error
}

// 缓存实体，包含一个缓存对象和一个结构通道
// 当数据还未加载到缓存时，在这个通道上接收数据将被阻塞
// 当数据已经加入到缓存后，这个通道将关闭
// 通道关闭后，任何在这个通道上接收数据将立刻得到一个零值，而不会阻塞线程
type entry struct {
	res   result
	ready chan struct{}
}

// 请求结构体，这个设计非常优雅，不愧是大师的手笔
// key是缓存的键值，response是一个缓存对象的通道
type request struct {
	key      int
	response chan<- result
}

// Memo 是缓存结构，我们在使用时就是用这个结构体来初始化
// requests通道巧妙的设计，使得可以使用独立线程来处理请求
// 后面的server方法就是专门用来处理请求
type Memo struct {
	requests chan request
}

// New 是初始化函数，当缓存不存在时，这个传入的函数用来向数据源取数据
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get 是获取缓存的方法，如果数据还未加入缓存，将用Func对应的函数取数据，然后存入缓存
func (memo *Memo) Get(key int) ([]string, error) {
	response := make(chan result)
	memo.requests <- request{key, response}

	// 接收通道，当数据还未载入缓存，将在这里阻塞，直到数据发送到这个通道
	res := <-response
	return res.value, res.err
}

// Close 是关闭请求
func (memo *Memo) Close() {
	close(memo.requests)
}

// 缓存数据请求的处理，这里有两个独立线程
// 一个是向数据源取数据，如果数据已经在缓存中，那么这个if内的代码将不会执行
// 一个是将缓存数据发送到request结构体的response通道，与Get方法的接受通道对应
func (memo *Memo) server(f Func) {
	cache := make(map[int]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

// 从数据源取数据，取完后关闭通道，于是这个通道将只允许接收数据，不允许发送数据
func (e *entry) call(f Func, key int) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

// 将缓存数据发送到request结构体的response通道
func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
