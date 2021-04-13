package rpcclient

import (
	"fmt"
	"github.com/duminghui/go-rpcclient/cmdjson"
	"testing"
)

func TestListReceivedByAddress(t *testing.T) {
	t.SkipNow()
	client := New(config)
	_, err := client.ListReceivedByAddress(0, false, false)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSendToAddress(t *testing.T) {
	t.SkipNow()
	client := New(config)
	//RSeTrvF3AR6L4Lwd5RF97QxeAybLaiokZ5
	tx, err := client.SendToAddress("4Lwd5RF97QxeAybLaiokZ5", 0.001)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tx)
}

func TestGetNewAddress(t *testing.T) {
	t.SkipNow()
	client := New(config)
	address, err := client.GetNewAddress("New")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("NewAddress", address)
}

func TestValidateAddress(t *testing.T) {
	t.SkipNow()
	client := New(config)
	client.Start()
	validateAddress, err := client.ValidateAddress("RRUwJogFBp9PVnrVs1RUHM1Xq728GNSz9asdfasdfsadfS")
	client.Shutdown()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", validateAddress)

}

func TestClient_ListUnspent(t *testing.T) {
	t.SkipNow()
	client := New(config)
	client.Start()
	unspentList, err := client.ListUnspent(nil, nil, nil)
	client.Shutdown()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("Size:%v\n", len(unspentList))
	//for _, unspent := range unspentList {
	//	fmt.Printf("unspect:%v\n", unspent)
	//}
}
func TestClient_CreateRawTransaction(t *testing.T) {
	t.SkipNow()
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

func TestClient_SignRawTransaction(t *testing.T) {
	hex := "010000000ba1f1257521f69dcf032efaa5c66307d70508494f4f5147080f0ea767971d8f160100000000ffffffffea3477cb54a8c51d06b26cf13fcda039b1a32086080d41bed6c4069a1822c8200100000000fffffffff17da5fd96aec7fa61bc17d8af5d44ad79f454cd6843881bd10f7237bbdb99290000000000ffffffff57bfab9c38abf49cfebd21103fad70050ef83a59bab1e8a2b1d049ece1af4a2f0000000000ffffffffca19fc8e444c4665a63abe45e7174a62a35536d72b2f577ff75c0c77b16d67370100000000fffffffffb4e0b1abaf75c73d34a419796522e6585f7f1334ed38dcb8aa2296f03c04d390000000000ffffffffeaeaa4b32dc7db7e57791b9c27622873d8a8784be3d9a4b011bab08f4bbc76470000000000ffffffff57783dc2f6b9eccbb89e6600eec5104f2a091a752aa443a8e9509386d66b65590100000000ffffffffff35c2dfea062abf673cecd4dc11da454054ed7c69d2d876bbf136537f0f715a0200000000ffffffffa878ea2a5e7fe3e87b431af1a81e99ddf1e3ab3a5634aeb4fd44770fa3b4135b0200000000ffffffffdd101f0651ba082244c48cf2c9480f377c08e640af00584ef93035fce2a3ab5d0200000000ffffffff0116254923000000001976a914fa529b8c14d146b373b7b55777b85135df15484788ac00000000"
	client := New(config)
	client.Start()
	txHex, err := client.SignRawTransaction(hex, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txHex)
}
