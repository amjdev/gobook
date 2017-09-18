package main

import "fmt"

func max(list ...int) int {
	m := 0
	for _, value := range list {
		if value > m {
			m = value
		}
	}
	return m
}

func main() {
	fmt.Println(max(1, 3, 5, 9, 6, 7, 8))
	fmt.Println(max(1, 30, 5, 9))
}
