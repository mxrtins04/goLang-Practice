package main

import "fmt"

func main() {
	var count int
	fmt.Print("How many numbers? ")
	fmt.Scan(&count)

	if count <= 0 {
		fmt.Println("Invalid count")
		return
	}

	var num, min, max int

	for i := 0; i < count; i++ {
		fmt.Print("Enter number: ")
		fmt.Scan(&num)

		if i == 0 {
			min = num
			max = num
		} else {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
	fmt.Println("Sum of extremes:", min+max)
}

