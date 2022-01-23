package FunctionalOptionPattern

import (
	"crypto/tls"
	"time"
)

//Server 要有侦听的 IP 地址 Addr 和端口号 Port ，这两个配置选项是必填的（当然，IP 地址和端口号都可以有默认值，
//不过这里我们用于举例，所以是没有默认值，而且不能为空，需要是必填的）。
//然后，还有协议 Protocol 、 Timeout 和MaxConns 字段， 这几个字段是不能为空的，但是有默认值的，比如，协议是 TCP，超时30秒 和 最大链接数1024个。
//还有一个 TLS ，这个是安全链接，需要配置相关的证书和私钥。这个是可以为空的
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

func NewDefaultServer(addr string, port int) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
	return &Server{addr, port, "tcp", timeout, 100, nil}, nil
}

func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil
}
