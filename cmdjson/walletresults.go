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

// GetTransactionDetailsResult models the details data from the gettransaction command.
//
// This models the "short" version of the ListTransactionsResult type, which
// excludes fields common to the transaction.  These common fields are instead
// part of the GetTransactionResult.
type GetTransactionDetailsResult struct {
	Account           string   `json:"account"`
	Address           string   `json:"address,omitempty"`
	Amount            float64  `json:"amount"`
	Category          string   `json:"category"`
	InvolvesWatchOnly bool     `json:"involveswatchonly,omitempty"`
	Fee               *float64 `json:"fee,omitempty"`
	Vout              uint32   `json:"vout"`
}

// GetTransactionResult models the data from the gettransaction command.
type GetTransactionResult struct {
	Amount          float64                       `json:"amount"`
	Fee             float64                       `json:"fee,omitempty"`
	Confirmations   int64                         `json:"confirmations"`
	BlockHash       string                        `json:"blockhash"`
	BlockIndex      int64                         `json:"blockindex"`
	BlockTime       int64                         `json:"blocktime"`
	TxID            string                        `json:"txid"`
	WalletConflicts []string                      `json:"walletconflicts"`
	Time            int64                         `json:"time"`
	TimeReceived    int64                         `json:"timereceived"`
	Details         []GetTransactionDetailsResult `json:"details"`
	Hex             string                        `json:"hex"`
}

// ListUnspentResult models a successful response from the listunspent request.
type ListUnspentResult struct {
	TxID          string  `json:"txid"`
	Vout          uint32  `json:"vout"`
	Address       string  `json:"address"`
	Account       string  `json:"account"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	RedeemScript  string  `json:"redeemScript,omitempty"`
	Amount        float64 `json:"amount"`
	Confirmations int64   `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
}
