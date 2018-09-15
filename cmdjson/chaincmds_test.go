package cmdjson

import (
	"fmt"
	"testing"
)

func TestGetConnectionCount(t *testing.T) {
	t.SkipNow()
	methods := RegisteredCmdMethods()
	fmt.Println(methods)
}
