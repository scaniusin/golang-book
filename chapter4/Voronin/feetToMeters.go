package main

import "fmt"

func main() {
	fmt.Print("Enter the number of feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(input, "feet = ", output, "meters")
}
