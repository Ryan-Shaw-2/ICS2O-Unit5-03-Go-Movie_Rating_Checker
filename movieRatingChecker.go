// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program checks what rating of movie the user can watch alone

package main

import "fmt"

func main() {
	// This function checks what rating of movie the user can watch alone
	var userAge int

	// input
	fmt.Println("This program checks what rating of movie the user can watch alone")
	fmt.Println()
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)
	fmt.Println()

	// process
	if userAge >= 17 {
		// output
		fmt.Println("You can go see an R rated movie alone")
	} else if userAge >= 13 {
		// output
		fmt.Println("You can go see a PG-13 rated movie alone")
	} else if userAge >= 5 {
		// output
		fmt.Println("You can go see a G or PG rated movie alone")
	} else {
		// output
		fmt.Println("You're too young")
	}
}
