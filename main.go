package main

import (
	"sync"

	"github.com/jerberlin/examples-ch9bank2/bank"
)

func main() {
	balance := bank.Balance()
	println("Initial balance: ", balance)

	var wg sync.WaitGroup

	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func() {
			// add some fixed amount
			amount := 1
			bank.Deposit(amount)
			wg.Done()
		}()
	}
	wg.Wait()

	// check final balance: should be 100
	balance = bank.Balance()
	println("Final balance: ", balance)
}
