package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBlanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBlanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	blanceText := fmt.Sprint(balance)
	os.WriteFile(accountBlanceFile, []byte(blanceText), 0644)
}

func main() {

	fmt.Println("Welcome to Go Bank!")
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("can't continue, sorry.")
	}

	for {

		presentOptions()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
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
			writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("thinks for choosing our bank")
			return
		}

	}

}
