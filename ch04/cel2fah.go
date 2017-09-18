package main

import "fmt"

func main() {
	var fahrenheit float64
	var celsius float64

	fmt.Print("Enter º fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Println(fahrenheit, "ºF = ", celsius, "ºC")
}
