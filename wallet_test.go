package rpcclient

import (
	"fmt"
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
	client := New(config)
	validateAddress, err := client.ValidateAddress("RRUwJogFBp9PVnrVs1RUHM1Xq728GNSz9S")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", validateAddress)
}
