package main

import "fmt"

func main() {
	var i int
	var m string
	for i = 1; i <= 100; i++ {
		m = ""
		if i%3 == 0 {
			m = "Fizz"
		}
		if i%5 == 0 {
			m += "Buzz"
		}

		if m == "" {
			fmt.Println(i)
		} else {
			fmt.Println(m)
		}
	}
}
