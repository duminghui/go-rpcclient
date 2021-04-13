// Package cmdjson provides ...
package cmdjson

type GetBlockCmd struct {
	Hash      string
	Verbose   *bool `jsonrpcdefault:"true"`
	VerboseTx *bool `jsonrpcdefault:"false"`
}

func NewGetBlockCmd(hash string, verbose, verboseTx *bool) *GetBlockCmd {
	return &GetBlockCmd{
		Hash:      hash,
		Verbose:   verbose,
		VerboseTx: verboseTx,
	}
}

type GetConnectionCountCmd struct{}

func NewGetConnectionCountCmd() *GetConnectionCountCmd {
	return &GetConnectionCountCmd{}
}

type ValidateAddressCmd struct {
	Address string
}

func NewValidateAddressCmd(address string) *ValidateAddressCmd {
	return &ValidateAddressCmd{
		Address: address,
	}
}

// TransactionInput represents the inputs to a transaction.  Specifically a
// transaction hash and output number pair.
type TransactionInput struct {
	Txid string `json:"txid"`
	Vout uint32 `json:"vout"`
}

// CreateRawTransactionCmd defines the createrawtransaction JSON-RPC command.
type CreateRawTransactionCmd struct {
	Inputs   []TransactionInput
	Amounts  map[string]float64 `jsonrpcusage:"{\"address\":amount,...}"` // In BTC
	LockTime *int64
}

// NewCreateRawTransactionCmd returns a new instance which can be used to issue
// a createrawtransaction JSON-RPC command.
//
// Amounts are in BTC. Passing in nil and the empty slice as inputs is equivalent,
// both gets interpreted as the empty slice.
func NewCreateRawTransactionCmd(inputs []TransactionInput, amounts map[string]float64,
	lockTime *int64) *CreateRawTransactionCmd {
	// to make sure we're serializing this to the empty list and not null, we
	// explicitly initialize the list
	if inputs == nil {
		inputs = []TransactionInput{}
	}
	return &CreateRawTransactionCmd{
		Inputs:   inputs,
		Amounts:  amounts,
		LockTime: lockTime,
	}
}

func init() {
	flags := UsageFlag(0)
	MustRegisterCmd("createrawtransaction", (*CreateRawTransactionCmd)(nil), flags)
	MustRegisterCmd("getconnectioncount", (*GetConnectionCountCmd)(nil), flags)
	MustRegisterCmd("getblock", (*GetBlockCmd)(nil), flags)
	MustRegisterCmd("validateaddress", (*ValidateAddressCmd)(nil), flags)
}
