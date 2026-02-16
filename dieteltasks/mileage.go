package main

import "fmt"

func main() {
    var miles, gallons int
    var totalMiles, totalGallons int

    for {
        fmt.Print("Enter miles driven (-1 to quit): ")
        fmt.Scan(&miles)
        if miles == -1 {
            break
        }

        fmt.Print("Enter gallons used: ")
        fmt.Scan(&gallons)

        tripMPG := float64(miles) / float64(gallons)
        fmt.Printf("Trip MPG: %.2f\n", tripMPG)

        totalMiles += miles
        totalGallons += gallons

        totalMPG := float64(totalMiles) / float64(totalGallons)
        fmt.Printf("Combined MPG so far: %.2f\n\n", totalMPG)
    }
}
