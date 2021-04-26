package netrcpcdemo2consumer

import (
	"fmt"
	"git/Uarealoser/rpcdemo/src/netrpc/netrpcdemo2"
	"net/rpc"
	"testing"
)

/**
安全一点的实现
*/
type HelloServiceInterface interface {
	Hello(request netrpcdemo2.User, reply *netrpcdemo2.User) error
}

type HelloServiceClient struct {
	*rpc.Client
}

func (c *HelloServiceClient) Hello(request netrpcdemo2.User, reply *netrpcdemo2.User) error {
	return c.Client.Call(netrpcdemo2.HelloServiceName+".Hello", request, reply)
}

func DialHelloService(address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{
		Client: client,
	}, nil
}

func TestConsumer1(t *testing.T) {
	client, err := DialHelloService("localhost:1234")
	if err != nil {
		t.Fatal("DialHelloService err:", err)
	}

	reply := netrpcdemo2.User{}
	if err := client.Hello(netrpcdemo2.User{
		Id:   1,
		Name: "hello",
		Role: netrpcdemo2.Consumer,
	}, &reply); err != nil {
		t.Fatal("client.Hello err:", err)
	}
	fmt.Println(reply)
}
