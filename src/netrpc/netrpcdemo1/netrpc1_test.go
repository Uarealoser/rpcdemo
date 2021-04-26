package netrpcdemo1

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

/**
基础实现
*/

type HelloService struct {
}

/*

 */
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func TestProvider(t *testing.T) {
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Fatal("ListenTCP err:", err)
	}

	conn, err := listen.Accept()
	if err != nil {
		t.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
	rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
}

func TestConsumer(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Fatal("dialing err:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		t.Fatal("rpc call err:", err)
	}
	fmt.Println(reply)
}
