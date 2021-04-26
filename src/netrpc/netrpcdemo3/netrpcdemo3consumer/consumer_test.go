package netrpcdemo3consumer

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestCounsumer(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatal("net.Dial err:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "ual", &reply)
	if err != nil {
		t.Fatal("client.Call err:", err)
	}
	fmt.Println(reply)

}
