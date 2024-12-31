// Exercise 2: Bank Transaction System

// Topics Covered: Go Constants, Go Loops, Go Break and Continue, Go Functions, Go
// Strings, Go Errors

// Case Study:

// You need to simulate a bank transaction system with the following features:

// 1. Account Management: Each account has an ID, name, and balance. Store the
// accounts in a slice.

// 2. Deposit Function: A function to deposit money into an account. Validate if the
// deposit amount is greater than zero.

// 3. Withdraw Function: A function to withdraw money from an account. Ensure the
// account has a sufficient balance before proceeding. Return appropriate errors
// for invalid amounts or insufficient balance.

// 4. Transaction History: Maintain a transaction history for each account as a string
// slice. Use a loop to display the transaction history when requested.

// 5. Menu System: Implement a menu-driven program where users can choose
// actions like deposit, withdraw, view balance, or exit. Use constants for menu
// options and break the loop to exit.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants for transaction types and menu options
const (
	DEPOSIT     = "DEPOSIT"
	WITHDRAWAL  = "WITHDRAWAL"
	MIN_BALANCE = 100.0 // Minimum balance required in account
)

// Custom errors
type BankError struct {
	Code    string
	Message string
}

func (e *BankError) Error() string {
	return fmt.Sprintf("Error [%s]: %s", e.Code, e.Message)
}

type Account struct {
	ID           int
	Name         string
	Balance      float64
	Transactions []Transaction
	CreatedAt    time.Time
}

type Transaction struct {
	Type      string
	Amount    float64
	Balance   float64
	Timestamp time.Time
}

// validateAmount checks if amount is valid for transactions
func validateAmount(amount float64) error {
	if amount <= 0 {
		return &BankError{
			Code:    "INVALID_AMOUNT",
			Message: "Amount must be greater than zero",
		}
	}
	return nil
}

// readInput reads and validates user input
func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return strings.TrimSpace(scanner.Text()), nil
}

// findAccount finds an account by ID
func findAccount(accounts []Account, id int) (int, error) {
	for i, acc := range accounts {
		if acc.ID == id {
			return i, nil
		}
	}
	return -1, &BankError{
		Code:    "ACCOUNT_NOT_FOUND",
		Message: "Account not found",
	}
}

func addAccount(accounts *[]Account) error {
	fmt.Println("\n=== Creating New Account ===")

	// Read ID
	idStr, err := readInput("Enter account ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &BankError{Code: "INVALID_ID", Message: "ID must be a number"}
	}

	// Check if ID exists
	if _, err := findAccount(*accounts, id); err == nil {
		return &BankError{Code: "DUPLICATE_ID", Message: "Account ID already exists"}
	}

	// Read name
	name, err := readInput("Enter account holder name: ")
	if err != nil {
		return err
	}
	if strings.TrimSpace(name) == "" {
		return &BankError{Code: "INVALID_NAME", Message: "Name cannot be empty"}
	}

	// Read initial balance
	balStr, err := readInput("Enter initial balance: ")
	if err != nil {
		return err
	}
	balance, err := strconv.ParseFloat(balStr, 64)
	if err != nil {
		return &BankError{Code: "INVALID_BALANCE", Message: "Balance must be a number"}
	}
	if balance < MIN_BALANCE {
		return &BankError{Code: "LOW_BALANCE", Message: fmt.Sprintf("Initial balance must be at least %.2f", MIN_BALANCE)}
	}

	// Create new account
	newAccount := Account{
		ID:        id,
		Name:      name,
		Balance:   balance,
		CreatedAt: time.Now(),
		Transactions: []Transaction{{
			Type:      "INITIAL_DEPOSIT",
			Amount:    balance,
			Balance:   balance,
			Timestamp: time.Now(),
		}},
	}

	*accounts = append(*accounts, newAccount)
	fmt.Printf("\n✅ Account created successfully for %s (ID: %d)\n", name, id)
	return nil
}

func deposit(accounts *[]Account) error {
	fmt.Println("\n=== Deposit Money ===")

	// Read account ID
	idStr, err := readInput("Enter account ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &BankError{Code: "INVALID_ID", Message: "ID must be a number"}
	}

	// Find account
	idx, err := findAccount(*accounts, id)
	if err != nil {
		return err
	}

	// Read amount
	amountStr, err := readInput("Enter deposit amount: ")
	if err != nil {
		return err
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return &BankError{Code: "INVALID_AMOUNT", Message: "Amount must be a number"}
	}
	if err := validateAmount(amount); err != nil {
		return err
	}

	// Process deposit
	(*accounts)[idx].Balance += amount
	(*accounts)[idx].Transactions = append((*accounts)[idx].Transactions, Transaction{
		Type:      DEPOSIT,
		Amount:    amount,
		Balance:   (*accounts)[idx].Balance,
		Timestamp: time.Now(),
	})

	fmt.Printf("\n✅ Successfully deposited %.2f. New balance: %.2f\n", amount, (*accounts)[idx].Balance)
	return nil
}

func withdraw(accounts *[]Account) error {
	fmt.Println("\n=== Withdraw Money ===")

	// Read account ID
	idStr, err := readInput("Enter account ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &BankError{Code: "INVALID_ID", Message: "ID must be a number"}
	}

	// Find account
	idx, err := findAccount(*accounts, id)
	if err != nil {
		return err
	}

	// Read amount
	amountStr, err := readInput("Enter withdrawal amount: ")
	if err != nil {
		return err
	}
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return &BankError{Code: "INVALID_AMOUNT", Message: "Amount must be a number"}
	}
	if err := validateAmount(amount); err != nil {
		return err
	}

	// Check sufficient balance
	if (*accounts)[idx].Balance-amount < MIN_BALANCE {
		return &BankError{
			Code:    "INSUFFICIENT_BALANCE",
			Message: fmt.Sprintf("Insufficient balance. Minimum balance required: %.2f", MIN_BALANCE),
		}
	}

	// Process withdrawal
	(*accounts)[idx].Balance -= amount
	(*accounts)[idx].Transactions = append((*accounts)[idx].Transactions, Transaction{
		Type:      WITHDRAWAL,
		Amount:    amount,
		Balance:   (*accounts)[idx].Balance,
		Timestamp: time.Now(),
	})

	fmt.Printf("\n✅ Successfully withdrew %.2f. New balance: %.2f\n", amount, (*accounts)[idx].Balance)
	return nil
}

func viewBalance(accounts []Account) error {
	fmt.Println("\n=== View Balance ===")

	// Read account ID
	idStr, err := readInput("Enter account ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &BankError{Code: "INVALID_ID", Message: "ID must be a number"}
	}

	// Find and display account details
	idx, err := findAccount(accounts, id)
	if err != nil {
		return err
	}

	acc := accounts[idx]
	fmt.Println("\n=== Account Details ===")
	fmt.Printf("Account Holder: %s\n", acc.Name)
	fmt.Printf("Account ID: %d\n", acc.ID)
	fmt.Printf("Current Balance: %.2f\n", acc.Balance)
	fmt.Printf("Account Created: %s\n", acc.CreatedAt.Format("2006-01-02 15:04:05"))
	return nil
}

func viewTransactionHistory(accounts []Account) error {
	fmt.Println("\n=== Transaction History ===")

	// Read account ID
	idStr, err := readInput("Enter account ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &BankError{Code: "INVALID_ID", Message: "ID must be a number"}
	}

	// Find account and display transactions
	idx, err := findAccount(accounts, id)
	if err != nil {
		return err
	}

	acc := accounts[idx]
	if len(acc.Transactions) == 0 {
		fmt.Println("No transactions found.")
		return nil
	}

	fmt.Printf("\nTransaction History for Account %d (%s)\n", acc.ID, acc.Name)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-20s %-12s %-15s %-15s\n", "Date", "Type", "Amount", "Balance")
	fmt.Println(strings.Repeat("-", 80))

	for _, t := range acc.Transactions {
		fmt.Printf("%-20s %-12s %-15.2f %-15.2f\n",
			t.Timestamp.Format("2006-01-02 15:04:05"),
			t.Type,
			t.Amount,
			t.Balance)
	}
	fmt.Println(strings.Repeat("-", 80))
	return nil
}
// insert sample data
func insertSampleData(accounts *[]Account) {
	*accounts = []Account{
		{ID: 1, Name: "John Doe", Balance: 1000, CreatedAt: time.Now()},
		{ID: 2, Name: "Jane Smith", Balance: 500, CreatedAt: time.Now()},
	}
}


func main() {
	accounts := []Account{}
	insertSampleData(&accounts)

	menuOptions := `
=== Bank Transaction System ===
1. Add Account
2. Deposit Money
3. Withdraw Money
4. View Balance
5. View Transaction History
6. Exit

Enter your choice (1-6): `

	for {
		choice, err := readInput(menuOptions)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		var operationErr error
		switch choice {
		case "1":
			operationErr = addAccount(&accounts)
		case "2":
			operationErr = deposit(&accounts)
		case "3":
			operationErr = withdraw(&accounts)
		case "4":
			operationErr = viewBalance(accounts)
		case "5":
			operationErr = viewTransactionHistory(accounts)
		case "6":
			fmt.Println("\n👋 Thank you for using our banking system. Goodbye!")
			return
		default:
			fmt.Println("\n❌ Invalid choice. Please try again.")
			continue
		}

		if operationErr != nil {
			fmt.Printf("\n❌ %v\n", operationErr)
		}
	}
}