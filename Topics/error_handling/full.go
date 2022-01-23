package error_handling

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var (
	// 长度不够，少一个Weight
	b = []byte{0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
	r = bytes.NewReader(b)
)

// Person is like Config in server.go
type Person struct {
	Name   [10]byte
	Age    uint8
	Weight uint8
	// nice !! holder of err
	err error
}

func (p *Person) read(data interface{}) {
	if p.err == nil {
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *Person) ReadName() *Person {
	p.read(&p.Name)
	return p
}

func (p *Person) ReadAge() *Person {
	p.read(&p.Age)
	return p
}

func (p *Person) ReadWeight() *Person {
	p.read(&p.Weight)
	return p
}

func (p *Person) Print() *Person {
	if p.err == nil {
		fmt.Printf("Name=%s, Age=%d, Weight=%d\n", p.Name, p.Age, p.Weight)
	}
	return p
}

// 相信你应该看懂这个技巧了，
// 不过，它的使用场景是有局限的，
// * 也就只能在对于同一个业务对象的不断操作下可以简化错误处理，
// 如果是多个业务对象，还是得需要各种 if err != nil的方式。

func Run() {
	p := Person{}

	p.ReadName().
		ReadAge().
		ReadWeight().
		Print()

	fmt.Println(p.err) // EOF 错误
}
