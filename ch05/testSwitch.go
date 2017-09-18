package main

import "fmt"

func main() {
	var i int
	for i = 1; i <= 100; i++ {
		switch (i%5) {
		//case 0: fmt.Println("Resto 0")
		case 0,1: fmt.Println("Resto 1")
		case 2: fmt.Println("Resto 2")
		case 3: fmt.Println("Resto 3")
		case 4: fmt.Println("Resto 4")	
		}
	}
}
