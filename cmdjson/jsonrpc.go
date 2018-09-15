// Package cmdjson provides ...
package cmdjson

import (
	"encoding/json"
	"fmt"
)

type RPCErrorCode int

type RPCError struct {
	Code    RPCErrorCode `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

func (e RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewPRCError(code RPCErrorCode, message string) *RPCError {
	return &RPCError{
		Code:    code,
		Message: message,
	}
}

func IsValidIDType(id interface{}) bool {
	switch id.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		string,
		nil:
		return true
	default:
		return false
	}
}

type RpcRequest struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      interface{}       `json:"id"`
}

func NewRpcRequest(id interface{}, method string, params []interface{}) (*RpcRequest, error) {
	if !IsValidIDType(id) {
		str := fmt.Sprintf("the id of type '%T' is invalid", id)
		return nil, makeError(ErrInvalidType, str)
	}
	rawParams := make([]json.RawMessage, 0, len(params))
	for _, param := range params {
		marshalledParam, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}
		rawMessage := json.RawMessage(marshalledParam)
		rawParams = append(rawParams, rawMessage)
	}
	return &RpcRequest{
		Jsonrpc: "1.0",
		Method:  method,
		Params:  rawParams,
		ID:      id,
	}, nil
}
