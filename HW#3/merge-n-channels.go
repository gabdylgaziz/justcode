package main

import (
	"fmt"
)

func mergeChannels(channels ...chan int) chan int {
	merged := make(chan int)
	
	for _, ch := range channels {
		go func(ch <-chan int) {
			for val := range ch {
				merged <- val
			}
		}(ch)
	}

	go func() {
		defer close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	merged := mergeChannels(ch1, ch2, ch3)

	for val := range merged {
		fmt.Println(val)
	}
}
