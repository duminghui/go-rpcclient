// Package cmdjson provides ...
package cmdjson

import (
	"fmt"
	"reflect"
)

func CmdMethod(cmd interface{}) (string, error) {
	rt := reflect.TypeOf(cmd)
	registerLock.RLock()
	method, ok := reflectTypeToMethod[rt]
	registerLock.RUnlock()
	if !ok {
		str := fmt.Sprintf("%q is not registerd", method)
		return "", makeError(ErrUnregisteredMethod, str)
	}
	return method, nil
}
