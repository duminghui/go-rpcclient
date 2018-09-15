package cmdjson

import (
	"fmt"
	"testing"
)

func TestMarshalCmd(t *testing.T) {
	jsoncmd, err := MarshalCmd("bb", NewGetBlockCmd("bb", nil, Bool(false)))
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%s\n", jsoncmd)
}
