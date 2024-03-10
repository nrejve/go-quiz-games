package main

import "fmt"

func screen() {
	fmt.Printf("\tWelcome to Aeronautics Quiz Game\n")
	fmt.Println("Enter your Name")
	fmt.Scanf("%s", &name)
}

func calc() {
	per = (float64(correct) / float64(3)) * 100

	if correct >= 2 {
		fmt.Println("You have passed")
	} else {
		fmt.Println("You have Failed")
	}

	fmt.Printf("Thanks %v !, Your details results are as follows\n", name)
	fmt.Printf("\t Total Questions: 3\n")
	fmt.Printf("\t Corrected Answers: %v\n", correct)
	fmt.Printf("\t Total Score: %v\n", score)
	fmt.Printf("\t Persantage of Score: %v\n", per)
}
