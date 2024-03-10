package main

import "fmt"

func questions() {
	fmt.Printf("\nQUES-1: Double decker Airbus twin engine Aircraft\n")
	fmt.Printf("\ta) A330\n")
	fmt.Printf("\tb) A380\n")
	fmt.Printf("\tc) A340\n")

	fmt.Println("Enter your answewr (a/b/c) bellow:")
	fmt.Scanf("%s", &answer)
	if answer == "b" {
		fmt.Println("Answer is correct")
		score = score + 5
		correct = correct + 1
	} else {
		fmt.Println("Answer is false")
	}

	fmt.Printf("\nQUES-2: Blackbox can have records of last\n")
	fmt.Printf("\ta) 1 Hour\n")
	fmt.Printf("\tb) 15 Minutes\n")
	fmt.Printf("\tc) 30 Minutes\n")

	fmt.Println("Enter your answewr (a/b/c) bellow:")
	fmt.Scanf("%s", &answer)
	if answer == "c" {
		fmt.Println("Answer is correct")
		score = score + 5
		correct = correct + 1
	} else {
		fmt.Println("Answer is false")
	}

	fmt.Printf("\nQUES-3: Airbus latest gen aircraft\n")
	fmt.Printf("\ta) Beluga\n")
	fmt.Printf("\tb) Neo\n")
	fmt.Printf("\tc) XWB\n")

	fmt.Println("Enter your answewr (a/b/c) bellow:")
	fmt.Scanf("%s", &answer)
	if answer == "b" {
		fmt.Println("Answer is correct")
		score = score + 5
		correct = correct + 1
	} else {
		fmt.Println("Answer is false")
	}
}
