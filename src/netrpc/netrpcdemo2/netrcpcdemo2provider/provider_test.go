package netrcpcdemo2provider

import (
	"git/Uarealoser/rpcdemo/src/netrpc/netrpcdemo2"
	"net"
	"net/rpc"
	"testing"
)

type HelloServiceInterface interface {
	Hello(request netrpcdemo2.User, reply *netrpcdemo2.User) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(netrpcdemo2.HelloServiceName, svc)
}

type HelloService struct {
}

func (p *HelloService) Hello(request netrpcdemo2.User, reply *netrpcdemo2.User) error {
	*reply = netrpcdemo2.User{
		Id:   request.Id,
		Name: request.Name,
		Role: netrpcdemo2.Provider,
	}
	return nil
}

func TestProvider(t *testing.T) {
	err := RegisterHelloService(new(HelloService))
	if err != nil {
		t.Fatal("RegisterHelloService err:", err)
	}
	listen, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		t.Fatal("net.Listen err:", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal("listen.Accept err:", err)
		}

		go rpc.ServeConn(conn)
	}
}
