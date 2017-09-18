package main

import "fmt"

func swap(x, y *int) {
	var temp int
	temp = *x
	// o bien... temp := *x
	*x = *y
	*y = temp
}

func main() {
	x := 1
	y := 2
	fmt.Println("X = ", x, " ; y = ", y)
	swap(&x, &y)
	fmt.Println("Swap!")
	fmt.Println("X = ", x, " ; y = ", y)
}
