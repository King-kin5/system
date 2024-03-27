package main

import (
	"bufio"
	"fmt"
	"os"
	"project-layout/account"
	"strconv"
	"strings"
)

func main() {
    myBank := account.NewBank()

	for {
		fmt.Print("Do you want to register a new account? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response != "yes" {
			break
		}

		fmt.Print("Enter your name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		newAccount := myBank.RegisterAccount(name)
		fmt.Println("Registration successful! Your account ID is:", newAccount.ID)
	}

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Transfer")
		fmt.Println("4. Inquiry")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)

		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter account ID: ")
			accountIDStr, _ := reader.ReadString('\n')
			accountID, _ := strconv.Atoi(strings.TrimSpace(accountIDStr))

			fmt.Print("Enter amount to deposit: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)

			err := myBank.Deposit(accountID, amount)
			if err != nil {
				fmt.Println("Deposit failed:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case 2:
			fmt.Print("Enter account ID: ")
			accountIDStr, _ := reader.ReadString('\n')
			accountID, _ := strconv.Atoi(strings.TrimSpace(accountIDStr))

			fmt.Print("Enter amount to withdraw: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)

			err := myBank.Withdraw(accountID, amount)
			if err != nil {
				fmt.Println("Withdrawal failed:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case 3:
			fmt.Print("Enter sender account ID: ")
			senderIDStr, _ := reader.ReadString('\n')
			senderID, _ := strconv.Atoi(strings.TrimSpace(senderIDStr))

			fmt.Print("Enter receiver account ID: ")
			receiverIDStr, _ := reader.ReadString('\n')
			receiverID, _ := strconv.Atoi(strings.TrimSpace(receiverIDStr))

			fmt.Print("Enter amount to transfer: ")
			amountStr, _ := reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)

			err := myBank.Transfer(senderID, receiverID, amount)
			if err != nil {
				fmt.Println("Transfer failed:", err)
			} else {
				fmt.Println("Transfer successful!")
			}

		case 4:
			fmt.Print("Enter account ID: ")
			accountIDStr, _ := reader.ReadString('\n')
			accountID, _ := strconv.Atoi(strings.TrimSpace(accountIDStr))

			acc, err := myBank.Inquiry(accountID)
			if err != nil {
				fmt.Println("Inquiry failed:", err)
			} else {
				fmt.Printf("Account ID: %d, Name: %s, Balance: %.2f\n", acc.ID, acc.Name, acc.Balance)
			}

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
    }

