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

type GetTransactionResult struct {
	Amount        float64                `json:"amount"`
	Fee           float64                `json:"fee,omitempty"`
	Confirmations int64                  `json:"confirmations"`
	BlockHash     string                 `json:"blockhash"`
	BlockIndex    int64                  `json:"blockindex"`
	BlockTime     int64                  `json:"blocktime"`
	Details       []GetTransactionDetail `json:"details"`
}

type GetTransactionDetail struct {
	Account  string  `json:"account"`
	Address  string  `json:"address"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Lable    string  `json:"lable"`
	Vout     uint32  `json:"vout"`
	Fee      float64 `json:"fee,omitempty"`
}
