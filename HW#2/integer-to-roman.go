package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	var b strings.Builder
	m := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	n := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < 13; i++ {
		j := num / m[i]
		num %= m[i]
		for j > 0 {
			b.WriteString(n[i])
			j--
		}
	}
	return b.String()
}

func main() {
	fmt.Println(intToRoman(555))
}
