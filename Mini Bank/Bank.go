package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func writeBalance(balance float64) {
	currentBalance := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(currentBalance), 0744)
}

func getBalanceFromFile() (float64, error){
	data,err := os.ReadFile(balanceFileName)
	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("Failed to parse the data in file")
	}
	return balance, nil
}

func main() {

	accountBalance, err := getBalanceFromFile()

	if(err != nil) {
		fmt.Println("Error: ", err)
		fmt.Println("=============")
		// panic("Can't continue")
	}

	fmt.Println("Welcome to Mini Bank!!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var option int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&option)

		switch option {
		case 1 :
			fmt.Println("Your account balance is: ", accountBalance)
		case 2: {
			var depositAmount float64
			fmt.Print("Enter amount to be deposited : ")
			fmt.Scan(&depositAmount)
			if depositAmount <=0 {
				fmt.Println("Invalid. Amount should be greater then 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Money Deposited: ", depositAmount)
			fmt.Println("Current balance: ", accountBalance)
			writeBalance(accountBalance)
		}
		case 3:
			var withdrawnAmount float64
			fmt.Print("Enter amount to be withdrawn : ")
			fmt.Scan(&withdrawnAmount)
			if withdrawnAmount <=0 {
				fmt.Println("Invalid. Amount should be greater then 0")
				continue
			}
			if withdrawnAmount > accountBalance {
				fmt.Println("Amount greater than account balance cannot be withdrawn!")
				continue
			}
			accountBalance -= withdrawnAmount
			fmt.Println("Money Withdrawn: ", withdrawnAmount)
			fmt.Println("Current balance: ", accountBalance)
			writeBalance(accountBalance)
		case 4:
			fmt.Println("GoodBye!! ")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
