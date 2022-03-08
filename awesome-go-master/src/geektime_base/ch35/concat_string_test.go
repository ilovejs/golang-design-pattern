package ch35

import (
	"bytes"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	assert := assert2.New(t)
	elements := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, element := range elements {
		ret += element
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByByteBuffer(t *testing.T) {
	assert := assert2.New(t)
	var buf bytes.Buffer
	elements := []string{"1", "2", "3", "4", "5"}
	for _, element := range elements {
		buf.WriteString(element)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elements := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, element := range elements {
			ret += element
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elements := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, element := range elements {
			buf.WriteString(element)
		}
	}
	b.StopTimer()
}
