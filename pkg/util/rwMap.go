package util

import "sync"

type RwMap struct {
	sync.RWMutex
	Store map[string]interface{}
}

func (m *RwMap) Set(key string, v interface{}) {
	m.Lock()
	m.Store[key] = v
	m.Unlock()
}

func (m *RwMap) Get(key string) interface{} {
	m.RLock()
	v := m.Store[key]
	m.RUnlock()
	return v
}

func NewRwMap() RwMap {
	return RwMap{
		RWMutex: sync.RWMutex{},
		Store:   make(map[string]interface{}),
	}
}
