package main

import "fmt"

func main() {
	var feet float64
	var meter float64

	fmt.Print("Enter feet: ")
	fmt.Scanf("%f", &feet)
	meter = feet * 0.3048
	fmt.Println(feet, " f = ", meter, " m")
}
