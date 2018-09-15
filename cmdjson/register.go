// Package cmdjson provides ...
package cmdjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
)

// UsageFlag usage flag
type UsageFlag uint32

const (
	// UFWalletOnly cmd only for had wallet
	UFWalletOnly UsageFlag = 1
)

type methodInfo struct {
	maxParams    int
	numReqParams int
	numOptParams int
	defaults     map[int]reflect.Value
	flags        UsageFlag
	usage        string
}

var (
	registerLock        sync.RWMutex
	methodToReflectType = make(map[string]reflect.Type)
	methodToInfo        = make(map[string]methodInfo)
	reflectTypeToMethod = make(map[reflect.Type]string)
)

func baseKindString(rt reflect.Type) string {
	numIndirects := 0
	for rt.Kind() == reflect.Ptr {
		numIndirects++
		rt = rt.Elem()
	}
	return fmt.Sprintf("%s%s", strings.Repeat("*", numIndirects), rt.Kind())
}

func isAcceptableKing(kind reflect.Kind) bool {
	switch kind {
	case reflect.Chan:
		fallthrough
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		return false
	}
	return true
}

func RegisterCmd(method string, cmd interface{}, flags UsageFlag) error {
	registerLock.Lock()
	defer registerLock.Unlock()
	if _, ok := methodToReflectType[method]; ok {
		str := fmt.Sprintf("method %q is already registered", method)
		return makeError(ErrDuplicateMethod, str)
	}
	if flags != 0 && flags != UFWalletOnly {
		str := fmt.Sprintf("invalid usage flags for method %s: %v", method, flags)
		return makeError(ErrInvalidUsageFlags, str)
	}
	rtp := reflect.TypeOf(cmd)
	if rtp.Kind() != reflect.Ptr {
		str := fmt.Sprintf("type must be *struct not '%s (%s)'", rtp, rtp.Kind())
		return makeError(ErrInvalidType, str)
	}
	rt := rtp.Elem()
	if rt.Kind() != reflect.Struct {
		str := fmt.Sprintf("type must be *struct not '%s (%s)'", rtp, rtp.Kind())
		return makeError(ErrInvalidType, str)
	}
	numFields := rt.NumField()
	numOptFields := 0
	defaults := make(map[int]reflect.Value)
	for i := 0; i < numFields; i++ {
		rtf := rt.Field(i)
		if rtf.Anonymous {
			str := fmt.Sprintf("embedded fields are not supported (field name: %s)", rtf.Name)
			return makeError(ErrEmbeddedType, str)
		}
		if rtf.PkgPath != "" {
			str := fmt.Sprintf("unexported files are not supported (field name: %q)", rtf.Name)
			return makeError(ErrUnexportedField, str)
		}
		var isOptional bool
		switch kind := rtf.Type.Kind(); kind {
		case reflect.Ptr:
			isOptional = true
			kind = rtf.Type.Elem().Kind()
			fallthrough
		default:
			if !isAcceptableKing(kind) {
				str := fmt.Sprintf("unsupported field type '%s (%s)' (field name %q)", rtf.Type, baseKindString(rtf.Type), rtf.Name)
				return makeError(ErrUnsupportedFieldType, str)
			}
		}

		if isOptional {
			numOptFields++
		} else {
			if numOptFields > 0 {
				str := fmt.Sprintf("all fields after the first option field must also be option (field name %q)", rtf.Name)
				return makeError(ErrNonOptionalField, str)
			}
		}

		if tag, ok := rtf.Tag.Lookup("jsonrpcdefault"); ok {
			if !isOptional {
				str := fmt.Sprintf("required fields must not have a default specified (field name %q)", rtf.Name)
				return makeError(ErrNonOptionalDefault, str)
			}
			rvf := reflect.New(rtf.Type.Elem())
			err := json.Unmarshal([]byte(tag), rvf.Interface())
			if err != nil {
				str := fmt.Sprintf("default value of %q is the wrong type (field name %q)", tag, rtf.Name)
				return makeError(ErrMismatchedDefault, str)
			}
			defaults[i] = rvf
		}
	}
	methodToReflectType[method] = rtp
	methodToInfo[method] = methodInfo{
		maxParams:    numFields,
		numReqParams: numFields - numOptFields,
		numOptParams: numOptFields,
		defaults:     defaults,
		flags:        flags,
		usage:        "",
	}
	reflectTypeToMethod[rtp] = method
	return nil
}

func MustRegisterCmd(method string, cmd interface{}, flags UsageFlag) {
	if err := RegisterCmd(method, cmd, flags); err != nil {
		panic(fmt.Sprintf("failed to register type %q: %v\n", method, err))
	}
}

func RegisteredCmdMethods() []string {
	registerLock.Lock()
	defer registerLock.Unlock()
	methods := make([]string, 0, len(methodToInfo))
	for k := range methodToInfo {
		methods = append(methods, k)
	}
	sort.Sort(sort.StringSlice(methods))
	return methods
}
