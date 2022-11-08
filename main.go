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

	//initstate.SetNonce()
	initstate.SetValue(0)
	if choice == 1 {

		contract.Deploy(initstate.Auth, initstate.Client)
	} else if choice == 2 {

		contract.CreateNewContract(initstate.Auth, initstate.Client)
	} else if choice == 3 {

		contract.GetBalance()
	} else if choice == 4 {

		initstate.SetValue(1e18)
		contract.ReceiverBalance(initstate.Auth)
	} else if choice == 5 {

		signer := common.HexToAddress("0xB273216C05A8c0D4F0a4Dd0d7Bae1D2EfFE636dd")
		to := common.HexToAddress("0x14723A09ACff6D2A60DcdF7aA4AFf308FDDC160C")
		amount := *big.NewInt(123)
		message := "coffee and donuts"
		signature := ([]byte)("0x993dab3dd91f5c6dc28e17439be475478f5635c92a56e17e82349d3fb2f166196f466c0b4e0c146f285204f0dcb13e5ae67bc33f4b888ec32dfe0a063e8f3f781b")
		nonce := *big.NewInt(1)
		contract.Signature(signer, to, &amount, message, &nonce, signature)
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