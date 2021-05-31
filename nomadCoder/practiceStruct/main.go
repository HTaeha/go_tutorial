package main

import (
	"fmt"
	accounts "go_tutorial/nomadCoder/practiceStruct/banking"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	account := accounts.NewAccount("taeha")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
