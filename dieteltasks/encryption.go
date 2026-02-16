package main

import "fmt"

func main() {
    var num int
    fmt.Print("Enter a 4-digit number: ")
    fmt.Scan(&num)

    d1 := (num / 1000) % 10
    d2 := (num / 100) % 10
    d3 := (num / 10) % 10
    d4 := num % 10

    d1 = (d1 + 7) % 10
    d2 = (d2 + 7) % 10
    d3 = (d3 + 7) % 10
    d4 = (d4 + 7) % 10

    d1, d3 = d3, d1
    d2, d4 = d4, d2

    encrypted := d1*1000 + d2*100 + d3*10 + d4
    fmt.Println("Encrypted:", encrypted)
}

