package main

import "fmt"

func half(n int) (h int, even bool) {
	if n%2 == 0 {
		return n / 2, true
	} else {
		return n / 2, false
	}
}

func main() {
	fmt.Println(half(11))
	fmt.Println(half(22))
}
