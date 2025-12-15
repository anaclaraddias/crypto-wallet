package main

import (
	"reflect"
	"testing"
)

func verifyResultExpectation(t *testing.T, expected, actual map[asset]float64) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected result was: %#v, but received: %#v", expected, actual)
	}
}

func TestExampleRun_Case1(t *testing.T) {
	currentWallet := map[asset]float64{}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
	}

	line := `{"type": "DEPOSIT", "asset": "BTC", "amount": 1.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_Case2(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 1.5,
	}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 1000,
	}

	line := `{"type": "DEPOSIT", "asset": "USD", "amount": 1000}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_Case3(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 1000,
	}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
	}

	line := `{"type": "WITHDRAW", "asset": "USD", "amount": 300}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_Case4(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
	}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
	}

	line := `{"type": "WITHDRAW", "asset": "BTC", "amount": 2.0}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_Case5(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
	}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
		"ETH": 5.0,
	}

	line := `{"type": "DEPOSIT", "asset": "ETH", "amount": 5.0}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_Case6(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 1.5,
		"USD": 700,
		"ETH": 5.0,
	}

	expectedWallet := map[asset]float64{
		"BTC": 1.0,
		"USD": 700,
		"ETH": 5.0,
	}

	line := `{"type": "WITHDRAW", "asset": "BTC", "amount": 0.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_InvalidType(t *testing.T) {
	currentWallet := map[asset]float64{}
	expectedWallet := map[asset]float64{}

	line := `{"type": "CREDIT", "asset": "BTC", "amount": 0.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_InvalidAsset(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 0.5,
	}

	expectedWallet := map[asset]float64{
		"BTC": 0.5,
	}

	line := `{"type": "DEPOSIT", "asset": "AAA", "amount": 0.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_InvalidAmount(t *testing.T) {
	currentWallet := map[asset]float64{
		"BTC": 0.5,
	}

	expectedWallet := map[asset]float64{
		"BTC": 0.5,
	}

	line := `{"type": "DEPOSIT", "asset": "ETH", "amount": -100}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}
