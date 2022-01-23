package with_config

import (
	"crypto/tls"
	"fmt"
	"time"
)

// NewServer 对于有洁癖的、有追求的程序员来说，他们会看到其中不太好的一点，那就是Config 并不是必需的，
// 所以，你需要判断是否是 nil 或是 Empty ——  Config{}会让我们的代码感觉不太干净

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

// Server should be just called Server, but since conflicts above
type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	// judge config here ...

	return &Server{
		Addr: "",
		Port: 0,
		Conf: nil,
	}, nil
}

func Run() {
	// default config
	srv1, _ := NewServer("localhost", 9000, nil)

	conf := Config{Protocol: "tcp", Timeout: 60 * time.Second}
	srv2, _ := NewServer("locahost", 9000, &conf)

	fmt.Println(srv1, srv2)
}
