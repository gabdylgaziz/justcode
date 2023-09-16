package main

import (
    "fmt"
)

func compareTwoSlices(a []int, b []int) bool{
    if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false;
		}
	}

	return true;
}

func main() {
    a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}

    
    fmt.Println(compareTwoSlices(a, b))
}

