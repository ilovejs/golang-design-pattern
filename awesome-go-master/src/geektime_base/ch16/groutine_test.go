package ch16

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	// 如果不加这行代码，可能上边还没有执行完，主程序就已经结束了
	time.Sleep(time.Millisecond * 50)
}
