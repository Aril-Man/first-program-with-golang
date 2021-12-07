package main

import "fmt"

func main() {
	fmt.Println("Walcome to my quiz game!")

	fmt.Printf("Enter your Name: ")

	var name string

	fmt.Scan(&name)

	fmt.Printf("Hello, %v ,Welcome to the game!\n", name)

	fmt.Printf("Enter your age : ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("Yay you can'y play!")
		return
	}

	score := 0
	num_questions := 3

	fmt.Printf("What is better, the RTX 3080 of RTX 3090? ")
	var answare string
	var answare2 string
	fmt.Scan(&answare, &answare2)

	if answare+" "+answare2 == "RTX 3090" || answare+" "+answare2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 5 1500x have? ")
	var core uint
	fmt.Scan(&core)

	if core == 6 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Your scored %v out of %v.\n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%.", percent)
}
