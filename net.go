// Package rpcclient provides ...
package rpcclient

import (
	"github.com/duminghui/go-rpcclient/cmdjson"
)

type FutureGetConnectionCountResult chan *serverResponse

func (r FutureGetConnectionCountResult) Receive() (int64, error) {
	resp, err := umarshalFuture(r, (*int64)(nil))
	if err != nil {
		return 0, err
	}
	return *(resp.(*int64)), nil
}

func (c *Client) GetConnectionCountAsync() FutureGetConnectionCountResult {
	cmd := cmdjson.NewGetConnectionCountCmd()
	return c.sendCmd(cmd)
}

func (c *Client) GetConnectionCount() (int64, error) {
	return c.GetConnectionCountAsync().Receive()
}
