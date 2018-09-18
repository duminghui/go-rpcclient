// Package rpcclient provides ...
package rpcclient

import "github.com/duminghui/go-rpcclient/cmdjson"

type FutureGetBlockResult chan *serverResponse

func (r FutureGetBlockResult) Receive() (string, error) {
	var result string
	err := umarshalFuture(r, &result)
	return result, err
}

func (c *Client) GetBlockAsync(blockHash string) FutureGetBlockResult {
	cmd := cmdjson.NewGetBlockCmd(blockHash, nil, cmdjson.Bool(false))
	return c.sendCmd(cmd)
}

func (c *Client) GetBlock(blockHash string) (string, error) {
	return c.GetBlockAsync(blockHash).Receive()
}
