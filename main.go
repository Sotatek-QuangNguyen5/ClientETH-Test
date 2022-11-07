package main

import (
	"fmt"

	"bufio"
	"os"
	"servercoin/contract"
	"servercoin/initstate"
)

var (

	in = bufio.NewReader(os.Stdin)
)

func handleChoice(choice int) {

	initstate.SetNonce()
	if choice == 1 {

		initstate.SetValue(0)
		contract.Deploy(initstate.Auth, initstate.Client)
	} else if choice == 2 {

		contract.CreateNewContract(initstate.Auth, initstate.Client)
	} else if choice == 3 {

		contract.GetBalance()
	} else if choice == 4 {

		initstate.SetValue(1e18)
		contract.ReceiverBalance(initstate.Auth)
	} else {

		fmt.Println("Input Invalid!!! Please enter again.")
	}
}

func main() {

	initstate.Init()
	for {

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Println("1 : Deploy Contract.")
		fmt.Println("2 : Create New Contract.")
		fmt.Println("3 : GetBalance Contract.")
		fmt.Println("4 : Send ETH to Contract.")
		fmt.Fscan(in, &choice)
		handleChoice(choice)
		fmt.Println()
	}
}