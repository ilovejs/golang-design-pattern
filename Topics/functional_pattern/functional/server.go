package functional

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// Core idea -- define an Option function use struct as input

type Option func(*Server)

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

// NewServer passing functions
func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {

	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}

	//...
	return &srv, nil
}

// Run 不但解决了“使用 Config 对象方式的需要有一个 config 参数，
// 但在不需要的时候，是放 nil 还是放 Config{}”的选择困难问题，
// 也不需要引用一个 Builder 的控制对象，直接使用函数式编程，在代码阅读上也很优雅
//	直觉式的编程；
//	高度的可配置化；
//	很容易维护和扩展；
//	自文档；
//	新来的人很容易上手；
//	没有什么令人困惑的事（是 nil 还是空）
func Run() {

	s1, _ := NewServer("localhost", 1024)

	s2, _ := NewServer("localhost", 2048,
		Protocol("udp"))

	s3, _ := NewServer("0.0.0.0", 8080,
		Timeout(300*time.Second),
		MaxConns(1000),
	)

	fmt.Println(s1, s2, s3)
}
