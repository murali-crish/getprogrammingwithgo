package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	var wearShades = true

	fmt.Println("You leave the cave:", exit)
	fmt.Println("Does command contains read", strings.Contains(command, "read"))
	fmt.Println("Wear Shades? ", wearShades)
}
