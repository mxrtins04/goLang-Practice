package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a non-negative integer: ")
    fmt.Scan(&n)

    factorial := 1

    for i := 1; i <= n; i++ {
        factorial *= i
    }

    fmt.Println("Factorial:", factorial)
}

