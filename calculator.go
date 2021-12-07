package main

import "fmt"

func main() {
	fmt.Println("Welcome to program calculator!")

	var name string
	fmt.Printf("Your Name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, Welcome\n", name)

	var operator string
	fmt.Printf("Select Operator (+ , - , * , / , **) : ")
	fmt.Scan(&operator)

	if operator == "+" || operator == "tambah" {
		var num1 uint
		var num2 uint
		fmt.Printf("Insert first number : ")
		fmt.Scan(&num1)
		fmt.Printf("Insert second number : ")
		fmt.Scan(&num2)

		prosess := (uint64(num1) + uint64(num2))

		fmt.Printf("Result of %v + %v = %v", num1, num2, prosess)
	} else if operator == "-" || operator == "kurang" {
		var num1 uint
		var num2 uint
		fmt.Printf("Insert first number : ")
		fmt.Scan(&num1)
		fmt.Printf("Insert second number : ")
		fmt.Scan(&num2)

		prosess := (uint64(num1) - uint64(num2))

		fmt.Printf("Result of %v - %v = %v", num1, num2, prosess)
	} else if operator == "*" || operator == "kali" {
		var num1 uint
		var num2 uint
		fmt.Printf("Insert first number : ")
		fmt.Scan(&num1)
		fmt.Printf("Insert second number : ")
		fmt.Scan(&num2)

		prosess := (uint64(num1) * uint64(num2))

		fmt.Printf("Result of %v * %v = %v", num1, num2, prosess)
	} else if operator == "/" || operator == "bagi" {
		var num1 uint
		var num2 uint
		fmt.Printf("Insert first number : ")
		fmt.Scan(&num1)
		fmt.Printf("Insert second number : ")
		fmt.Scan(&num2)

		prosess := (uint64(num1) / uint64(num2))

		fmt.Printf("Result of %v / %v = %v", num1, num2, prosess)
	} else if operator == "**" || operator == "pangkat" {
		var num1 uint
		var num2 uint
		fmt.Printf("Insert first number : ")
		fmt.Scan(&num1)
		fmt.Printf("Insert second number : ")
		fmt.Scan(&num2)

		prosess := (uint64(num1) / uint64(num2))

		fmt.Printf("Result of %v ** %v = %v", num1, num2, prosess)
	} else {
		fmt.Println("Operator not found!")
	}

}
