package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "˚ F\n")
	fmt.Print((9.0/5.0*celsius)+32, "˚ F\n")
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "˚ F\n")
	piggyBank := 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
