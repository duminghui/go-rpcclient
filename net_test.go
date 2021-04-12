package rpcclient

import (
	"fmt"
	"testing"
)

var config = &ConnConfig{
	Name:    "BCRM",
	Host:    "127.0.0.1:12095",
	User:    "tmp",
	Pass:    "tmp",
	LogJSON: true,
}

func TestGetConnectionCount(t *testing.T) {
	//t.SkipNow()
	client := New(config)
	client.Start()
	//count, err := client.GetConnectionCount()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(count)
	//count, err = client.GetConnectionCount()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(count)
	str, err := client.GetBlock("00001dee87831d24dd3e4b7b911e634a7d613c08a3dd2bbc32ee2f8cd8f66dbe")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(str)
}

func TestGetConnectionCount2(t *testing.T) {
	t.SkipNow()
	client := New(config)
	client.Start()
}
