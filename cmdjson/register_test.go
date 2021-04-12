package cmdjson

import (
	"fmt"
	"testing"
)

func TestRegisteredCmdMethods(t *testing.T) {
	methods := RegisteredCmdMethods()
	for _, method := range methods {
		fmt.Println(method)
	}
}
