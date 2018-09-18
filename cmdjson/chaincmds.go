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

func init() {
	MustRegisterCmd("getconnectioncount", (*GetConnectionCountCmd)(nil))
	MustRegisterCmd("getblock", (*GetBlockCmd)(nil))
	MustRegisterCmd("validateaddress", (*ValidateAddressCmd)(nil))
}
