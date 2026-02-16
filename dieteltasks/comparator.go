package main

import "fmt"

func main() {
    var x1, y1, x2, y2 int

    fmt.Println("Enter x1 y1:")
    fmt.Scan(&x1, &y1)

    fmt.Println("Enter x2 y2:")
    fmt.Scan(&x2, &y2)

    if x1 == x2 {
        fmt.Println("Line is perpendicular to the x-axis (vertical line).")
    } else if y1 == y2 {
        fmt.Println("Line is perpendicular to the y-axis (horizontal line).")
    } else {
        fmt.Println("Line is not perpendicular to any axis.")
    }
}

