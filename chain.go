// Package rpcclient provides ...
package rpcclient

import "github.com/duminghui/go-rpcclient/cmdjson"

type FuntureGetBlockResult chan *serverResponse

func (r FuntureGetBlockResult) Receive() (string, error) {
	resp, err := umarshalFuture(r, (*string)(nil))
	if err != nil {
		return "", err
	}
	return *(resp.(*string)), nil
}

func (c *Client) GetBlockAsync(blockHash string) FuntureGetBlockResult {
	cmd := cmdjson.NewGetBlockCmd(blockHash, nil, cmdjson.Bool(false))
	return c.sendCmd(cmd)
}

func (c *Client) GetBlock(blockHash string) (string, error) {
	return c.GetBlockAsync(blockHash).Receive()
}
