// Package cmdjson provides ...
package cmdjson

type ListReceivedByAddressResult struct {
	Account       string   `json:"account"`
	Address       string   `json:"address"`
	Amount        float64  `json:"amount"`
	Confirmations uint64   `json:"confirmations"`
	TxIDs         []string `json:"txids"`
}

type ValidateAddressResult struct {
	IsValid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	IsMine       bool   `json:"ismine"`
	IsWatchOnly  bool   `json:"iswatchonly"`
	IsScript     bool   `json:"isScript"`
	IsCompressed bool   `json:"iscompressed"`
}
