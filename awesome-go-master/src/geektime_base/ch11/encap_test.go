package ch11

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

//func (e *Employee) string() string {
//	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name: %s/Age:%d", e.Id, e.Name, e.Age)
//}
func (e Employee) string() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObjj(t *testing.T) {
	e := Employee{"0", "Bob", 10}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Name = "Jack"
	e2.Age = 22
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Name)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.string())
}
