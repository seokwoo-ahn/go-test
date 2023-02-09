package map_test

import "sync"

type MapWithMutex struct {
	store map[int]int
	mu    sync.RWMutex
}

func NewMapWithMutex() *MapWithMutex {
	return &MapWithMutex{store: make(map[int]int)}
}

func (mw *MapWithMutex) Set(key, val int) {
	mw.mu.Lock()
	defer mw.mu.Unlock()
	mw.store[key] = val
}

func (mw *MapWithMutex) Get(key int) (int, bool) {
	mw.mu.RLock()
	defer mw.mu.RUnlock()
	val, ok := mw.store[key]
	return val, ok
}

type MapWithOutMutex struct {
	store map[int]int
}

func NewMapWithOutMutex() *MapWithOutMutex {
	return &MapWithOutMutex{store: make(map[int]int)}
}

func (m *MapWithOutMutex) Set(key, val int) {
	m.store[key] = val
}

func (m *MapWithOutMutex) Get(key int) (int, bool) {
	val, ok := m.store[key]
	return val, ok
}
