package main

import "fmt"

func main() {
    var population float64
    var rate float64

    fmt.Print("Enter current population: ")
    fmt.Scan(&population)

    fmt.Print("Enter annual growth rate (%): ")
    fmt.Scan(&rate)

    rate = rate / 100

    fmt.Println("Year\tPopulation\tIncrease")

    for year := 1; year <= 75; year++ {
        increase := population * rate
        population += increase
        fmt.Printf("%d\t%.0f\t%.0f\n", year, population, increase)
    }
}

