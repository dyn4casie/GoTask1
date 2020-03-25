package main

import "fmt"

func main() {
	var input int

	fmt.Print("Enter a decimal value to see it equivalent binary value: ")
	fmt.Scan(&input)

	hex := fmt.Sprintf("%b", input)
	fmt.Printf("Decimal %d => Binary: %s\n", input, hex)
}
