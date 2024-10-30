package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/abrahamkifleT/Bank/fileops"
	"github.com/dustin/go-humanize"
)

const accountBlanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBlanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 phone: ", randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			formatAccountBalance := humanize.Commaf(accountBalance)
			fmt.Println("Your balance is: ", formatAccountBalance)
		case 2:
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
			fileops.WriteFloatToFile(accountBalance, accountBlanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			} else if withdrawAmount > accountBalance {
				fmt.Println("Your balance is inceffcent")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("successfully withdraw: ", withdrawAmount)
				fmt.Println("Your current balance: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBlanceFile)
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("thinks for choosing our bank")
			return
		}

	}

}
