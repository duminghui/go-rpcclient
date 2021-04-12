package cmdjson

import (
	"fmt"
	"testing"
)

func TestMarshalCmd(t *testing.T) {
	t.SkipNow()
	jsoncmd, err := MarshalCmd(RpcVersion1, "this is id", NewGetBlockCmd("bb", nil, Bool(false)))
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%s\n", jsoncmd)
}

func TestNewCmd(t *testing.T) {
	tmp, err := NewCmd("getblock", "1000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmp)
}
