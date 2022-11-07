package readfile

import (
    "bufio"
    "log"
    "os"
)

func ReadContractOne() (string) {


    f, err := os.Open("./text/contract.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var contractAddress string

    for scanner.Scan() {

		contractAddress = scanner.Text()
		break
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return contractAddress
}

func ReadContractTwo() (string) {


    f, err := os.Open("./text/contract.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var contractAddress string

    for scanner.Scan() {

		contractAddress = scanner.Text()
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return contractAddress
}

func ReadPrivateKey() (string) {


    f, err := os.Open("./text/privatekey.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
	var privateKey string

    for scanner.Scan() {

		privateKey = scanner.Text()
		break
    }

    if err := scanner.Err(); err != nil {

        log.Fatal(err)
    }

	return privateKey
}