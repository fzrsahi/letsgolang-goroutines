package letsgolang

import (
	"fmt"
	"sync"
	"testing"
)

type BankAccount struct {
	Balance int
	RwMutex sync.RWMutex
}

func (account *BankAccount) AddBalance(amount int) {
	account.RwMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RwMutex.Lock()
	balance := account.Balance
	account.RwMutex.Unlock()
	return balance
}

func TestRwMutex(t *testing.T) {
	account := BankAccount{}

	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func(group *sync.WaitGroup) {
			defer group.Done()
			group.Add(1)
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}(group)
	}

	group.Wait()

	fmt.Println("Final Balance :", account.GetBalance())

}
