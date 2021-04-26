package netrpcdemo4provider

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

/**
基于Http协议的RPC：
RPC的服务架设在"/jsonrpc"路径，在处理函数中基于Http.ResponseWriter和http.Request类型的参数构造
一个io.ReadWriterCloser通道。然后基于conn构建针对服务端的json编码解码器，最后通过rpc.ServeRequest
函数为每次请求处理一次RPC方法调用
调用：curl localhost:1234/jsonrpc -X POST \
    --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
*/
func TestProvider(t *testing.T) {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
