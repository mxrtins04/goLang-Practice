package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum:", sum)
}

