package main

import "fmt"

func main() {
	var room = "lake"

	switch room {
	case "cave":
		fmt.Println("You find yourself in a dim lit cavern.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freesing cold.")
	}
}
