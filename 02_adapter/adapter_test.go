package adapter

import "testing"

var expect = "adaptee method"

// this test is 'Client'
func TestAdapter(t *testing.T) {

	adaptee := NewAdaptee()

	// wrapping !
	// Adapter 是转换Adaptee为Target接口的适配器
	target := NewAdapter(adaptee)

	// client use target interface method
	res := target.Request()

	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
