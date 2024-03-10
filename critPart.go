package main

import "fmt"

func critPart() {
	round = 1
	for {
		score = 0
		correct = 0
		screen()
		subscriptions()
		questions()
		calc()

		fmt.Println("Want to continue this quiz again? (yes/no)")
		fmt.Scanf("%s", &answer)
		if answer == "yes" {
			round = round + 1
			fmt.Printf("\tStarting Round Number %v\n", round)
			continue
		} else if answer == "no" {
			fmt.Printf("Thanks %v for pariticipate in this quiz game, hope will see you in future !\n", name)
			fmt.Printf("You have completed %v rounds\n", round)
			break
		}

	}
}
