package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// The file in which we cheep info
const accountBalanceFile = "balance.txt"

// Extract data from file
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Fail to find the file!")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Fail to parse the store data!")
	}

	return balance, nil
}

// Add data in file
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	//use data from file
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Print("________________")
		return
		//panic("Can't continue, sorry!")
	}

	fmt.Println("Welcome to go Bank!")

	//for i := 0; i < 2; i++ {
	for {
		fmt.Println("What if you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice here: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greter that 0.")
				//return
				continue
			}
			accountBalance = accountBalance + depositAmount
			//accountBalance += depositAmount
			fmt.Println("Balance update! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greter that 0.")
				//return
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withraw more taht you have.")
				//return
				continue
			}
			accountBalance = accountBalance - withdrawAmount
			//accountBalance += depositAmount
			fmt.Println("Balance update! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
