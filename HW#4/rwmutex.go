package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	value int
	mu    sync.RWMutex
}

func (d *Data) Read() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.value
}

func (d *Data) Write(newValue int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.value = newValue
}

func main() {
	data := &Data{value: 0}

	for i := 0; i < 5; i++ {
		go func() {
			for {
				value := data.Read()
				fmt.Println("Read:", value)
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	go func() {
		for i := 1; i <= 10; i++ {
			data.Write(i)
			fmt.Println("Write:", i)
			time.Sleep(time.Second)
		}
	}()

	select {}
}
