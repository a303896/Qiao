package bank

import "sync"

var balance int

var mu sync.RWMutex

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func deposit(amount int) {
	balance = balance + amount
}

func Deposit2(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance2() int {
	mu.RLock()
	//defer mu.RUnlock()
	return balance
}
