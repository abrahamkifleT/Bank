package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to Go Bank!")
	var accountBalance = 1000.0

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)
			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater then 0.")
				//return
				continue
			}
			accountBalance += depositeAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
			} else if withdrawAmount > accountBalance {
				fmt.Println("Your balance is inceffcent")
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("successfully withdraw: ", withdrawAmount)
				fmt.Println("Your current balance: ", accountBalance)
			}
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("thinks for choosing our bank")

}
