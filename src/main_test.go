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

func TestExampleRun_DepositBTC(t *testing.T) {
	currentWallet := map[asset]float64{}

	expectedWallet := map[asset]float64{
		"BTC": 1.5,
	}

	line := `{"type": "DEPOSIT", "asset": "BTC", "amount": 1.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestExampleRun_DepositUSD(t *testing.T) {
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

func TestExampleRun_WithdrawUSD(t *testing.T) {
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

func TestExampleRun_InvalidWithdrawBTC(t *testing.T) {
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

func TestExampleRun_DepositETH(t *testing.T) {
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

func TestExampleRun_WithdrawBTC(t *testing.T) {
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

func TestInvalidType(t *testing.T) {
	currentWallet := map[asset]float64{}
	expectedWallet := map[asset]float64{}

	line := `{"type": "CREDIT", "asset": "BTC", "amount": 0.5}`
	walletResult := processTransaction(line, currentWallet)

	verifyResultExpectation(t, expectedWallet, walletResult)
}

func TestInvalidAsset(t *testing.T) {
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

func TestInvalidAmount(t *testing.T) {
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
