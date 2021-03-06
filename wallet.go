// Package rpcclient provides ...
package rpcclient

import (
	"github.com/duminghui/go-rpcclient/cmdjson"
)

type FutureListReceivedByAddressResult chan *serverResponse

func (r FutureListReceivedByAddressResult) Receive() ([]cmdjson.ListReceivedByAddressResult, error) {
	var result []cmdjson.ListReceivedByAddressResult
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) ListReceivedByAddress(minConf int, includeEmpty, includeWatchOnly bool) ([]cmdjson.ListReceivedByAddressResult, error) {
	return c.ListReceivedByAddressAsync(minConf, includeEmpty, includeWatchOnly).Receive()
}

func (c *Client) ListReceivedByAddressAsync(minConf int, includeEmpty, includeWatchOnly bool) FutureListReceivedByAddressResult {
	cmd := cmdjson.NewListReceivedByAddressCmd(&minConf, &includeEmpty, &includeWatchOnly)
	return c.sendCmd(cmd)
}

type FutureSendToAddressResult chan *serverResponse

func (r FutureSendToAddressResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) SendToAddressAsync(address string, amount float64) FutureSendToAddressResult {
	cmd := cmdjson.NewSendToAddressCmd(address, amount)
	return c.sendCmd(cmd)
}

func (c *Client) SendToAddress(address string, amount float64) (string, error) {
	return c.SendToAddressAsync(address, amount).Receive()
}

type FutureGetNewAddressResult chan *serverResponse

func (r FutureGetNewAddressResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) GetNewAddressAsync(account string) FutureGetNewAddressResult {
	cmd := cmdjson.NewGetNewAddressCmd(&account)
	return c.sendCmd(cmd)
}

func (c *Client) GetNewAddress(account string) (string, error) {
	return c.GetNewAddressAsync(account).Receive()
}

type FutureValidateAddressResult chan *serverResponse

func (r FutureValidateAddressResult) Receive() (*cmdjson.ValidateAddressResult, error) {
	var result cmdjson.ValidateAddressResult
	err := unmarshalFuture(r, &result)
	return &result, err
}

func (c *Client) ValidateAddressAsync(address string) FutureValidateAddressResult {
	cmd := cmdjson.NewValidateAddressCmd(address)
	return c.sendCmd(cmd)
}

func (c *Client) ValidateAddress(address string) (*cmdjson.ValidateAddressResult, error) {
	return c.ValidateAddressAsync(address).Receive()
}

type FutureGetTransactionResult chan *serverResponse

func (r FutureGetTransactionResult) Receive() (*cmdjson.GetTransactionResult, error) {
	var result cmdjson.GetTransactionResult
	err := unmarshalFuture(r, &result)
	return &result, err
}

func (c *Client) GetTransactionAsync(txid string, includeWatchOnly *bool) FutureGetTransactionResult {
	cmd := cmdjson.NewGetTransactionCmd(txid, includeWatchOnly)
	return c.sendCmd(cmd)
}

func (c *Client) GetTransaction(txid string, includeWatchOnly *bool) (*cmdjson.GetTransactionResult, error) {
	return c.GetTransactionAsync(txid, includeWatchOnly).Receive()
}

type FutureListUnspentResult chan *serverResponse

func (r FutureListUnspentResult) Receive() ([]cmdjson.ListUnspentResult, error) {
	var result []cmdjson.ListUnspentResult
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) ListUnspentAsync(minConf, maxConf *int, addrs *[]string) FutureListUnspentResult {
	cmd := cmdjson.NewListUnspentCmd(minConf, maxConf, addrs)
	return c.sendCmd(cmd)
}

func (c *Client) ListUnspent(minConf, maxConf *int, addr *[]string) ([]cmdjson.ListUnspentResult, error) {
	return c.ListUnspentAsync(minConf, maxConf, addr).Receive()
}

type CreateRawTransactionResult chan *serverResponse

func (r CreateRawTransactionResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) CreateRawTransactionAsync(inputs []cmdjson.TransactionInput, amounts map[string]float64, lockTime *int64) CreateRawTransactionResult {
	cmd := cmdjson.NewCreateRawTransactionCmd(inputs, amounts, lockTime)
	return c.sendCmd(cmd)
}

func (c *Client) CreateRawTransaction(inputs []cmdjson.TransactionInput, amounts map[string]float64, lockTime *int64) (string, error) {
	return c.CreateRawTransactionAsync(inputs, amounts, lockTime).Receive()
}

type SignRawTransactionResult chan *serverResponse

func (r SignRawTransactionResult) Receive() (cmdjson.SignRawTransactionResult, error) {
	var result cmdjson.SignRawTransactionResult
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) SignRawTransactionAsync(hexEncodedTx string, inputs *[]cmdjson.RawTxInput, privKeys *[]string, flags *string) SignRawTransactionResult {
	cmd := cmdjson.NewSignRawTransactionCmd(hexEncodedTx, inputs, privKeys, flags)
	return c.sendCmd(cmd)
}

func (c *Client) SignRawTransaction(hexEncodedTx string, inputs *[]cmdjson.RawTxInput, privKeys *[]string, flags *string) (cmdjson.SignRawTransactionResult, error) {
	return c.SignRawTransactionAsync(hexEncodedTx, inputs, privKeys, flags).Receive()
}
