// Package rpcclient provides ...
package rpcclient

import (
	"github.com/duminghui/go-rpcclient/cmdjson"
)

type FutureGetBlockResult chan *serverResponse

func (r FutureGetBlockResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) GetBlockAsync(blockHash string) FutureGetBlockResult {
	cmd := cmdjson.NewGetBlockCmd(blockHash, nil, cmdjson.Bool(false))
	return c.sendCmd(cmd)
}

func (c *Client) GetBlock(blockHash string) (string, error) {
	return c.GetBlockAsync(blockHash).Receive()
}

type FutureSendRawTransactionResult chan *serverResponse

func (r FutureSendRawTransactionResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) SendRawTransactionAsync(hexTx string, allowHighFees *bool) FutureSendRawTransactionResult {
	cmd := cmdjson.NewSendRawTransactionCmd(hexTx, allowHighFees)
	return c.sendCmd(cmd)
}

// SendRawTransaction submits the encoded transaction to the server which will
// then relay it to the network.
func (c *Client) SendRawTransaction(hexTx string, allowHighFees *bool) (string, error) {
	return c.SendRawTransactionAsync(hexTx, allowHighFees).Receive()
}
