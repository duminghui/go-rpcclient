// Package rpcclient provides ...
package rpcclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/duminghui/go-rpcclient/cmdjson"
	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidAuth    = errors.New("authentication failure")
	ErrClientShutdown = errors.New("the client has been shutdown")
)

const (
	sendPostBufferSize = 100
)

var log = logrus.New()

func SetLog(logger *logrus.Logger) {
	log = logger
}

type rpcResponse struct {
	Result json.RawMessage   `json:"result"`
	Error  *cmdjson.RPCError `json:"error"`
}

type sendPostDetails struct {
	httpRequest   *http.Request
	serverRequest *serverRequest
}

type serverRequest struct {
	id                 uint64
	method             string
	cmd                interface{}
	marshalledJSON     []byte
	serverResponseChan chan *serverResponse
}

// serverResponse is the raw byte of a JSON-RPC result
// or the error if the serverResponse error object was non-null
type serverResponse struct {
	result []byte
	err    error
}

func (r rpcResponse) result() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}

// Client to connection RPC server
type Client struct {
	id           uint64 // cmd id
	config       *ConnConfig
	httpClient   *http.Client
	sendPostChan chan *sendPostDetails
	shutdown     chan struct{}
	wg           sync.WaitGroup
}

// NextID next id for JSON-RPC message
func (c *Client) NextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

func newFutureError(err error) chan *serverResponse {
	serverResponseChan := make(chan *serverResponse, 1)
	serverResponseChan <- &serverResponse{err: err}
	return serverResponseChan
}

func umarshalFuture(c chan *serverResponse, resp interface{}) error {
	r := <-c
	respBytes, err := r.result, r.err
	if err != nil {
		return err
	}
	return json.Unmarshal(respBytes, resp)
}

func receiveFuture(c chan *serverResponse) ([]byte, error) {
	r := <-c
	return r.result, r.err
}

func (c *Client) handleSendPostMessage(details *sendPostDetails) {
	serverReq := details.serverRequest
	log.Infof("[%s]Sending post message [%s] with id %d", c.config.Name, serverReq.method, serverReq.id)
	httpResp, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		serverReq.serverResponseChan <- &serverResponse{err: err}
		return
	}
	respBytes, err := ioutil.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		serverReq.serverResponseChan <- &serverResponse{err: err}
		return
	}
	if c.config.LogJSON {
		log.Infof("[%s]Sending command content [%s](%d)\n%s", c.config.Name, serverReq.method, serverReq.id, serverReq.marshalledJSON)
		var indentJSONOut bytes.Buffer
		err = json.Indent(&indentJSONOut, respBytes, "", "  ")
		if err != nil {
			log.Errorln("Indent response json error")
		} else {
			log.Infof("[%s]RPC response [%s](%d)\n%s", c.config.Name, serverReq.method, serverReq.id, indentJSONOut.Bytes())
		}
	}
	var resp rpcResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		err = fmt.Errorf("status code: %d, response: %q", httpResp.StatusCode, string(respBytes))
		serverReq.serverResponseChan <- &serverResponse{err: err}
		return
	}
	res, err := resp.result()
	serverReq.serverResponseChan <- &serverResponse{result: res, err: err}
}

func (c *Client) sendPostHandler() {
	func() {
		for {
			select {
			case details := <-c.sendPostChan:
				c.handleSendPostMessage(details)
			case <-c.shutdown:
				return
			}
		}
	}()

	func() {
		for {
			select {
			case details := <-c.sendPostChan:
				details.serverRequest.serverResponseChan <- &serverResponse{
					result: nil,
					err:    ErrClientShutdown,
				}
			default:
				return
			}
		}
	}()
	c.wg.Done()
	log.Infof("[%s]RPC client send handler done for %s", c.config.Name, c.config.Host)
}

func (c *Client) sendPostRequest(httpReq *http.Request, serverReq *serverRequest) {
	select {
	case <-c.shutdown:
		serverReq.serverResponseChan <- &serverResponse{result: nil, err: ErrClientShutdown}
	default:
	}
	c.sendPostChan <- &sendPostDetails{
		serverRequest: serverReq,
		httpRequest:   httpReq,
	}
}

func (c *Client) sendPost(serverReq *serverRequest) {
	url := "http://" + c.config.Host
	bodyReader := bytes.NewReader(serverReq.marshalledJSON)
	httpReq, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		serverReq.serverResponseChan <- &serverResponse{result: nil, err: err}
		return
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(c.config.User, c.config.Pass)
	log.Infof("[%s]Sending command [%s] with id %d", c.config.Name, serverReq.method, serverReq.id)
	c.sendPostRequest(httpReq, serverReq)
}

func (c *Client) sendCmd(cmd interface{}) chan *serverResponse {
	method, err := cmdjson.CmdMethod(cmd)
	if err != nil {
		return newFutureError(err)
	}
	id := c.NextID()
	marshalledJSON, err := cmdjson.MarshalCmd(id, cmd)
	if err != nil {
		return newFutureError(err)
	}
	serverResponseChan := make(chan *serverResponse, 1)
	serverReq := &serverRequest{
		id:                 id,
		method:             method,
		cmd:                cmd,
		marshalledJSON:     marshalledJSON,
		serverResponseChan: serverResponseChan,
	}
	c.sendPost(serverReq)
	return serverResponseChan
}

func (c *Client) doShutdown() bool {
	select {
	case <-c.shutdown:
		return false
	default:
	}
	log.Infof("[%s]Shutting down RPC client %s", c.config.Name, c.config.Host)
	close(c.shutdown)
	return true
}

func (c *Client) Shutdown() {
	if !c.doShutdown() {
		return
	}
}
func (c *Client) start() {
	log.Infof("[%s]Starting RPC client %s", c.config.Name, c.config.Host)
	c.wg.Add(1)
	go c.sendPostHandler()
}

func (c *Client) WaitForShutdown() {
	c.wg.Wait()
}

// ConnConfig RPC server conn info
type ConnConfig struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	LogJSON bool   `json:"logJSON"`
}

func NewConfig(confFile string) (*ConnConfig, error) {
	configBytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	var connConfig ConnConfig
	err = json.Unmarshal(configBytes, &connConfig)
	return &connConfig, err
}

func New(config *ConnConfig) *Client {
	client := &Client{
		config:       config,
		httpClient:   &http.Client{},
		sendPostChan: make(chan *sendPostDetails, sendPostBufferSize),
		shutdown:     make(chan struct{}),
	}
	client.start()
	return client
}
