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
	//cmd, err := cmdjson.NewCmd("getblock", blockHash)
	//if err != nil {
	//	fmt.Println(err)
	//	panic("DDDDD")
	//
	//}
	return c.sendCmd(cmd)
}

func (c *Client) GetBlock(blockHash string) (string, error) {
	return c.GetBlockAsync(blockHash).Receive()
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
