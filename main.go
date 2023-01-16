package main

import (
	"fmt"

	banking "github.com/Esppe0w/goproject_1/bankaccount"
)

func main() {
	account := banking.Account{Owner: "희재", Balance: 1000}

	fmt.Println(account)
}
