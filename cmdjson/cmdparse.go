// Package cmdjson provides ...
package cmdjson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func makeParams(rt reflect.Type, rv reflect.Value) []interface{} {
	numFileds := rt.NumField()
	params := make([]interface{}, 0, numFileds)
	for i := 0; i < numFileds; i++ {
		rft := rt.Field(i)
		rvf := rv.Field(i)
		if rft.Type.Kind() == reflect.Ptr {
			if rvf.IsNil() {
				break
			}
			// TODO why do this
			rvf.Elem()
		}
		params = append(params, rvf.Interface())
	}

	return params
}

func MarshalCmd(id interface{}, cmd interface{}) ([]byte, error) {
	rt := reflect.TypeOf(cmd)
	registerLock.RLock()
	method, ok := reflectTypeToMethod[rt]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registered", method)
		return nil, makeError(ErrUnregisteredMethod, str)
	}
	rv := reflect.ValueOf(cmd)
	if rv.IsNil() {
		str := "the specified command is nil"
		return nil, makeError(ErrInvalidType, str)
	}
	params := makeParams(rt.Elem(), rv.Elem())
	rawCmd, err := NewRpcRequest(id, method, params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(rawCmd)
}
