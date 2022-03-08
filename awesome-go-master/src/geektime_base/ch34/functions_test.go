package ch34

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	type args struct {
		op int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"the arg is 2", args{2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.op); got != tt.want {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	//t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	//t.Fatal("Error")
	fmt.Println("End")
}

func TestSquare2WithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		assert.Equal(t, Square(inputs[i]), expected[i])
	}
}
