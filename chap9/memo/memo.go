package memo

import "sync"

// F is the type of the function to memoize.
type F func(string) (interface{}, error)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     F
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type result struct {
	value interface{}
	err   error
}

// New returns a instance of memo
func New(f F) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

// Get returns content fetching by given url
func (m *Memo) Get(url string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[url]
	if e == nil {
		// First time this entry gets hit
		// fetch content, construct entry & broadcast
		e = &entry{ready: make(chan struct{})}
		m.cache[url] = e
		m.mu.Unlock()
		e.res.value, e.res.err = m.f(url)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
