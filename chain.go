// Package rpcclient provides ...
package rpcclient

import (
	"fmt"
	"github.com/duminghui/go-rpcclient/cmdjson"
)

type FutureGetBlockResult chan *serverResponse

func (r FutureGetBlockResult) Receive() (string, error) {
	var result string
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) GetBlockAsync(blockHash string) FutureGetBlockResult {
	//cmd := cmdjson.NewGetBlockCmd(blockHash, nil, cmdjson.Bool(false))
	cmd, err := cmdjson.NewCmd("getblock", blockHash)
	if err != nil {
		fmt.Println(err)
		panic("DDDDD")

	}
	return c.sendCmd(cmd)
}

func (c *Client) GetBlock(blockHash string) (string, error) {
	return c.GetBlockAsync(blockHash).Receive()
}
