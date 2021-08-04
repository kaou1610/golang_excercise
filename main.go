package main

import (
	"fmt"
	
)

func calculate(o string, a, b float64) {
	switch o {
	case "add": 
		fmt.Println("a + b = ", a + b)
	case "sub":
		fmt.Println("a - b = ", a - b)
	case "mul":
		fmt.Println("a*b = ", a*b)
	case "div":
		if b == 0 {
			fmt.Print("Error")
		} else {
			fmt.Println("a/b = ", a/b)
		}
		
	}
}

func main()  {
	fmt.Println("Hello world!")
	calculate("mul", 5, 4)
}