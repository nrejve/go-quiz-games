package main

import (
	"fmt"
	"os"
)

func subscriptions() {
	count = 3

	for {
		fmt.Printf("\tEnter the 9 digit subscription code to play(i.e XXX-XXX-XXX): ")

		fmt.Scanf("%s", &sc)

		if sc == "B76-EWD-NR9" {
			fmt.Println("Worked !")
			break
		} else if sc == "EK4-REN-0PB" {
			fmt.Println("Worked !")
			break
		} else if sc == "T6N-9JS-AFB" {
			fmt.Println("Worked !")
			break
		} else {
			fmt.Printf("Not worked ! Try Again (%v/3 Attempts Left)\n", count)
			count = count - 1
			if count == 0 {
				fmt.Println("\nReached Maximum Tries, EXITING ...")
				os.Exit(1)
			}
		}
	}
}
