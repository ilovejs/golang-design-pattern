package builder

import (
	"crypto/tls"
	"fmt"
	"time"
)

/** In java,
User user = new User.Builder()
  .name("Hao Chen")
  .email("haoel@hotmail.com")
  .nickname("左耳朵")
  .build();
*/

type Server struct {
	Addr string
	Port int
	// optional
	Protocol string
	MaxConns int

	Timeout time.Duration
	TLS     *tls.Config
}

//ServerBuilder 使用一个builder类来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	//其它代码设置其它成员的默认值
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithMaxConn(maxconn int) *ServerBuilder {
	sb.Server.MaxConns = maxconn
	return sb
}

func (sb *ServerBuilder) WithTimeOut(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() Server {
	return sb.Server
}

// Run 这种方式也很清楚，不需要额外的 Config 类，使用链式的函数调用的方式来构造一个对象，
//只需要多加一个 Builder 类。你可能会觉得，这个 Builder 类似乎有点多余，
// * 我们似乎可以直接在Server 上进行这样的 Builder 构造，的确是这样的。
//   但是，在处理错误的时候可能就有点麻烦，不如一个包装类更好一些
func Run() {

	sb := ServerBuilder{}

	server := sb.Create("127.0.0.1", 8080).
		WithProtocol("udp").
		WithMaxConn(1024).
		WithTimeOut(30 * time.Second).
		Build()

	// there is no err handling

	fmt.Println(server)
}
