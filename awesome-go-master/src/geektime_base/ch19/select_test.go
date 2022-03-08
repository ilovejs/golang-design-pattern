package ch19

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "Service Down"
}

func AsyncService() chan string {
	retch := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("return result.")
		retch <- ret
		fmt.Println("Service exited.")
	}()
	return retch
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
