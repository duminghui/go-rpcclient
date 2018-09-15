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

func init() {
	flags := UsageFlag(0)
	MustRegisterCmd("getconnectioncount", (*GetConnectionCountCmd)(nil), flags)
	MustRegisterCmd("getblock", (*GetBlockCmd)(nil), flags)
}
