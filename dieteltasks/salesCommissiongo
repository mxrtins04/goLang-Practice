package main

import "fmt"

func main() {
    var itemNumber int
    var totalSales float64

    fmt.Println("Enter item number (1-4) sold last week. Enter 0 to finish:")

    for {
        fmt.Print("Item: ")
        fmt.Scan(&itemNumber)

        if itemNumber == 0 {
            break
        }

        if itemNumber == 1 {
            totalSales += 239.99
        } else if itemNumber == 2 {
            totalSales += 129.75
        } else if itemNumber == 3 {
            totalSales += 99.95
        } else if itemNumber == 4 {
            totalSales += 350.89
        } else {
            fmt.Println("Invalid item number.")
        }
    }

    earnings := 200.0 + (totalSales * 0.09)

    fmt.Printf("Total sales: $%.2f\n", totalSales)
    fmt.Printf("Salesperson's earnings: $%.2f\n", earnings)
}
