// Package cmdjson provides ...
package cmdjson

type ListReceivedByAddressCmd struct {
	MinConf          *int  `jsonrpcdefault:"1"`
	IncludeEmpty     *bool `jsonrpcdefault:"false"`
	IncludeWatchOnly *bool `jsonrpcdefault:"false"`
}

func NewListReceivedByAddressCmd(minConf *int, includeEmpty, includeWatchOnly *bool) *ListReceivedByAddressCmd {
	return &ListReceivedByAddressCmd{
		MinConf:          minConf,
		IncludeEmpty:     includeEmpty,
		IncludeWatchOnly: includeWatchOnly,
	}
}

type SendToAddressCmd struct {
	Address string
	Amount  float64
}

func NewSendToAddressCmd(address string, amount float64) *SendToAddressCmd {
	return &SendToAddressCmd{
		Address: address,
		Amount:  amount,
	}
}

type GetNewAddressCmd struct {
	Account *string
}

func NewGetNewAddressCmd(account *string) *GetNewAddressCmd {
	return &GetNewAddressCmd{
		Account: account,
	}
}

type GetTransactionCmd struct {
	TxID             string
	IncludeWatchOnly *bool
}

func NewGetTransactionCmd(txID string, includeWathcOnly *bool) *GetTransactionCmd {
	return &GetTransactionCmd{
		TxID:             txID,
		IncludeWatchOnly: includeWathcOnly,
	}
}

// ListUnspentCmd defines the listunspent JSON-RPC command.
type ListUnspentCmd struct {
	MinConf   *int `jsonrpcdefault:"1"`
	MaxConf   *int `jsonrpcdefault:"9999999"`
	Addresses *[]string
}

// NewListUnspentCmd returns a new instance which can be used to issue a
// listunspent JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewListUnspentCmd(minConf, maxConf *int, addresses *[]string) *ListUnspentCmd {
	return &ListUnspentCmd{
		MinConf:   minConf,
		MaxConf:   maxConf,
		Addresses: addresses,
	}
}

// RawTxInput models the data needed for raw transaction input that is used in
// the SignRawTransactionCmd struct.
type RawTxInput struct {
	Txid         string `json:"txid"`
	Vout         uint32 `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
	RedeemScript string `json:"redeemScript"`
}

// SignRawTransactionCmd defines the signrawtransaction JSON-RPC command.
type SignRawTransactionCmd struct {
	RawTx    string
	Inputs   *[]RawTxInput
	PrivKeys *[]string
	Flags    *string `jsonrpcdefault:"\"ALL\""`
}

// NewSignRawTransactionCmd returns a new instance which can be used to issue a
// signrawtransaction JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewSignRawTransactionCmd(hexEncodedTx string, inputs *[]RawTxInput, privKeys *[]string, flags *string) *SignRawTransactionCmd {
	return &SignRawTransactionCmd{
		RawTx:    hexEncodedTx,
		Inputs:   inputs,
		PrivKeys: privKeys,
		Flags:    flags,
	}
}

func init() {
	flags := UFWalletOnly
	MustRegisterCmd("getnewaddress", (*GetNewAddressCmd)(nil), flags)
	MustRegisterCmd("gettransaction", (*GetTransactionCmd)(nil), flags)
	MustRegisterCmd("listreceivedbyaddress", (*ListReceivedByAddressCmd)(nil), flags)
	MustRegisterCmd("listunspent", (*ListUnspentCmd)(nil), flags)
	MustRegisterCmd("sendtoaddress", (*SendToAddressCmd)(nil), flags)
	MustRegisterCmd("signrawtransaction", (*SignRawTransactionCmd)(nil), flags)
}
