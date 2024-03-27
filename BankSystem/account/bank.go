package account
import (
    "errors"
)

// Bank manages bank operations
type Bank struct {
    accounts map[int]*Account
}

// NewBank initializes a new Bank instance
func NewBank() *Bank {
    return &Bank{
        accounts: make(map[int]*Account),
    }
}

// PutAccount adds or updates an account in the bank
func (b *Bank) PutAccount(account *Account) {
    b.accounts[account.ID] = account
}

// GetAccount retrieves an account from the bank based on account ID
func (b *Bank) GetAccount(id int) (*Account, bool) {
    account, ok := b.accounts[id]
    return account, ok
}

// RegisterAccount registers a new account in the bank
func (b *Bank) RegisterAccount(name string) *Account {
    // Generate a unique account ID (you can use a simple counter for demonstration purposes)
    // You may need to handle concurrency if this application is used by multiple users simultaneously
    accountID := generateUniqueID()

    // Create an account object with the provided details
    newAccount := &Account{
        ID:      accountID,
        Name:    name,
        Balance: 0, // Initial balance is usually zero for a new account
    }

    // Store the new account information in the bank
    b.PutAccount(newAccount)

    return newAccount
}

// Deposit deposits an amount to the account with given ID
func (b *Bank) Deposit(accountID int, amount float64) error {
    if amount < 0 {
        return errors.New("amount must be positive")
    }

    account, ok := b.GetAccount(accountID)
    if !ok {
        return errors.New("account not found")
    }

    account.Balance += amount
    b.PutAccount(account)

    return nil
}

// Withdraw withdraws an amount from the account with given ID
func (b *Bank) Withdraw(accountID int, amount float64) error {
    if amount < 0 {
        return errors.New("amount must be positive")
    }

    account, ok := b.GetAccount(accountID)
    if !ok {
        return errors.New("account not found")
    }

    if account.Balance < amount {
        return errors.New("insufficient balance")
    }

    account.Balance -= amount
    b.PutAccount(account)

    return nil
}

// Transfer transfers an amount from one account to another
func (b *Bank) Transfer(senderID, receiverID int, amount float64) error {
    if amount < 0 {
        return errors.New("amount must be positive")
    }

    sender, ok := b.GetAccount(senderID)
    if !ok {
        return errors.New("sender account not found")
    }

    receiver, ok := b.GetAccount(receiverID)
    if !ok {
        return errors.New("receiver account not found")
    }

    if sender.Balance < amount {
        return errors.New("insufficient balance")
    }

    sender.Balance -= amount
    receiver.Balance += amount

    b.PutAccount(sender)
    b.PutAccount(receiver)

    return nil
}

// Inquiry retrieves account information by ID
func (b *Bank) Inquiry(accountID int) (*Account, error) {
    account, ok := b.GetAccount(accountID)
    if !ok {
        return nil, errors.New("account not found")
    }
    return account, nil
}

func generateUniqueID() int {
    // Implement your logic to generate a unique ID here
    return 0 // Placeholder value, replace with actual logic
}
