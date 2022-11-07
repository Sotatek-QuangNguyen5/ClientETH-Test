solc:
	solc --abi --bin --overwrite contracts/ExchangeCoin.sol -o build

abigen:
	abigen --bin=build/VerifySignature.bin --abi=build/VerifySignature.abi --pkg=exchangecoin --out=exchangecoin/ExchangeCoin.go

run:
	go run main.go