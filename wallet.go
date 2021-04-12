// Package rpcclient provides ...
package rpcclient

import "github.com/duminghui/go-rpcclient/cmdjson"

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
