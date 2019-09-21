package main

import "fmt"

func main() {
	var year = 2100
	if year%4 == 0 || (year%4 == 0 && year%100 > 0) {
		fmt.Printf("Year %v is leap year\n", year)
	} else {
		fmt.Printf("Year %v is not leap year\n", year)
	}
}
