package main

import (
	"fmt"
	"sync"
)

type MySyncMap struct {
	data map[string]interface{}
	mu   sync.Mutex
}

func NewMySyncMap() *MySyncMap {
	return &MySyncMap{
		data: make(map[string]interface{}),
	}
}

func (m *MySyncMap) Store(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *MySyncMap) Load(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *MySyncMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, key)
}

func main() {
	syncMap := NewMySyncMap()

	syncMap.Store("key1", 42)
	syncMap.Store("key2", "Hello, World!")

	value, found := syncMap.Load("key1")
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("not found")
	}
}
