package main

import (
	"fmt"
	"sort"
)

type SliceWithSort []int

func (s SliceWithSort) Sort(){
	sort.Ints(s)
}

func main() {
	s := SliceWithSort{5, 2, 9, 1, 5, 6}
	s.Sort()
	fmt.Println(s)
}
