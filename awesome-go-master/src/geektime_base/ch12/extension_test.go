package ch12

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

/**
总结就是在 GO语言中，无法实现和 java 中的逻辑一样，LSP 原则，不能够重写父类的某个方法
*/
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("chao")
}
