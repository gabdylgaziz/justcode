package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println(value)
	}()

	go func() {
		ch <- 42
	}()

	select {}
}
