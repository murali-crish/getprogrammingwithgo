package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)
	fmt.Printf("Type %T for %[1]v\n", "murali")
	fmt.Printf("Type %T for %[1]v\n", 3.14)
	fmt.Printf("Type %T for %[1]v\n", 42)
	fmt.Printf("Type %T for %[1]v\n", true)

	var red, green, blue uint8 = 0, 141, 213
	fmt.Printf("%x %x %x\n", red, green, blue)
	fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)

	var yellow uint8 = 255
	yellow++
	fmt.Println(yellow)

	var number int8 = 127
	number++
	fmt.Println(number)
}
