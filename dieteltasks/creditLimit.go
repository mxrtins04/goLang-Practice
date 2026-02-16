package main

import "fmt"

func main() {
    var accountNum, beginBalance, charges, credits, limit int

    fmt.Print("Enter account number: ")
    fmt.Scan(&accountNum)
    fmt.Print("Enter beginning balance: ")
    fmt.Scan(&beginBalance)
    fmt.Print("Enter total charges: ")
    fmt.Scan(&charges)
    fmt.Print("Enter total credits: ")
    fmt.Scan(&credits)
    fmt.Print("Enter credit limit: ")
    fmt.Scan(&limit)

    newBalance := beginBalance + charges - credits

    fmt.Printf("New balance for account %d is: %d\n", accountNum, newBalance)

    if newBalance > limit {
        fmt.Println("Credit limit exceeded")
    }
}
