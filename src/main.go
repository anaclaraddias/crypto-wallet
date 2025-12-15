package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	wallet := map[asset]float64{}

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if line == "stop\n" {
			break
		}

		transactionResult := processTransaction(line, wallet)

		jsonBytes, err := json.Marshal(transactionResult)
		if err != nil {
			fmt.Printf("error while parsing string into json: %v\n", err)
			break
		}

		fmt.Println(string(jsonBytes))
	}
}

func processTransaction(line string, wallet map[asset]float64) map[asset]float64 {
	var transaction Transaction
	if err := json.Unmarshal([]byte(line), &transaction); err != nil {
		fmt.Printf("error while parsing json: %v\n", err)
		return wallet
	}

	if !isTransactionValid(transaction) {
		return wallet
	}

	if transaction.Type == TypeDeposit {
		wallet[transaction.Asset] += transaction.Amount
		return wallet
	}

	if transaction.Amount > wallet[transaction.Asset] {
		fmt.Printf("attention: withdraw amount need to be less or equal than the wallet asset amount\n\n")
		return wallet
	}

	wallet[transaction.Asset] -= transaction.Amount

	return wallet
}

func isTransactionValid(transaction Transaction) bool {
	if !slices.Contains(ValidAssets, transaction.Asset) {
		fmt.Printf("attention: invalid asset\n\n")
		return false
	}

	if !slices.Contains(ValidTypes, transaction.Type) {
		fmt.Printf("attention: invalid type\n\n")
		return false
	}

	if transaction.Amount < 0 {
		fmt.Printf("attention: invalid amount\n\n")
		return false
	}

	return true
}
