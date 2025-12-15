package main

type Transaction struct {
	Asset  asset           `json:"asset"`
	Type   transactionType `json:"type"`
	Amount float64         `json:"amount"`
}

type asset string
type transactionType string

const (
	AssetBTC asset = "BTC"
	AssetETH asset = "ETH"
	AssetUSD asset = "USD"

	TypeWithdraw transactionType = "WITHDRAW"
	TypeDeposit  transactionType = "DEPOSIT"
)

var ValidAssets = []asset{
	AssetBTC,
	AssetETH,
	AssetUSD,
}

var ValidTypes = []transactionType{
	TypeWithdraw,
	TypeDeposit,
}
