package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBlanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}
func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {

	fmt.Println("Welcome to Go Bank!")
	var accountBalance, err = getFloatFromFile(accountBlanceFile)
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
			writeFloatToFile(accountBalance, accountBlanceFile)
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
				writeFloatToFile(accountBalance, accountBlanceFile)
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("thinks for choosing our bank")
			return
		}

	}

}
