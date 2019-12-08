package main

import "fmt"

func main() {
	for _, i := range []int{1, 2, 3, 4, 5} {
		for _, j := range []int{6, 7, 8, 9, 10} {
			if i == 2 {
				break
			}
			fmt.Println(j)
		}
		fmt.Println(i)
	}
}
