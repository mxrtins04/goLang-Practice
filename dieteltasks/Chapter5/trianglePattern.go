package main

import "fmt"

func main() {

	for index := 1; index <= 10; index++ {
		for j := 0; j < index; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()


	for index := 10; index >= 1; index-- {
		for x := 0; x < index; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	for index := 1; index <= 10; index++ {
		for x1 := 0; x1 < 10-index; x1++ {
			fmt.Print(" ")
		}
		for x2 := 0; x2 < index; x2++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()


	for index := 10; index >= 1; index-- {
		for x3 := 0; x3 < 10-index; x3++ {
			fmt.Print(" ")
		}
		for x4 := 0; x4 < index; x4++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

