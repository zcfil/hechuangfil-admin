package eth

import (
	"io"
	"net/http"
)

//http客户端接口
type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

//日志接口
type logger interface {
	Println(v ...interface{})
}

// WithHttpClient set custom http client
func WithHttpClient(client httpClient) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.client = client
	}
}

// WithLogger set custom logger
func WithLogger(l logger) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.log = l
	}
}

// WithDebug set debug flag
func WithDebug(enabled bool) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.Debug = enabled
	}
}
