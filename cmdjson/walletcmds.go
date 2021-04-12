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
func init() {
	flags := UFWalletOnly
	MustRegisterCmd("listreceivedbyaddress", (*ListReceivedByAddressCmd)(nil), flags)
	MustRegisterCmd("sendtoaddress", (*SendToAddressCmd)(nil), flags)
	MustRegisterCmd("getnewaddress", (*GetNewAddressCmd)(nil), flags)
	MustRegisterCmd("gettransaction", (*GetTransactionCmd)(nil), flags)
}
