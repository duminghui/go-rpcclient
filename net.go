// Package rpcclient provides ...
package rpcclient

import (
	"github.com/duminghui/go-rpcclient/cmdjson"
)

type FutureGetConnectionCountResult chan *serverResponse

func (r FutureGetConnectionCountResult) Receive() (int64, error) {
	var result int64
	err := unmarshalFuture(r, &result)
	return result, err
}

func (c *Client) GetConnectionCountAsync() FutureGetConnectionCountResult {
	cmd := cmdjson.NewGetConnectionCountCmd()
	return c.sendCmd(cmd)
}

func (c *Client) GetConnectionCount() (int64, error) {
	return c.GetConnectionCountAsync().Receive()
}
