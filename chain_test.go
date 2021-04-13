package rpcclient

import (
	"fmt"
	"github.com/duminghui/go-rpcclient/cmdjson"
	"testing"
)

func TestClient_CreateRawTransaction(t *testing.T) {
	tmp := cmdjson.ListUnspentResult{}
	tmp2 := cmdjson.TransactionInput{
		Txid: tmp.TxID,
		Vout: tmp.Vout,
	}
	client := New(config)
	client.Start()
	txHex, err := client.CreateRawTransaction([]cmdjson.TransactionInput{tmp2}, map[string]float64{"1111111111111111111": 0.01}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txHex)
}
