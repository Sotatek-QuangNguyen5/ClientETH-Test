package main

import (
	"fmt"
	"math/big"

	"bufio"
	"os"
	"servercoin/contract"
	"servercoin/initstate"

	"github.com/ethereum/go-ethereum/common"
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
	} else if choice == 5 {

		initstate.SetValue(0)
		signer := common.HexToAddress("0x9EB46ba1179259642706f03387ADC39470f64Cc8")
		to := common.HexToAddress("0x9EB46ba1179259642706f03387ADC39470f64Cc8")
		amount := *big.NewInt(1e18)
		message := "oke"
		signature := []byte("0x9EB46ba1179259642706f03387ADC39470f64Cc8")
		contract.Signature(signer, to, &amount, message, initstate.Auth.Nonce, signature)
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
		fmt.Println("5 : Signature ETH to Contract.")
		fmt.Fscan(in, &choice)
		handleChoice(choice)
		fmt.Println()
	}
}