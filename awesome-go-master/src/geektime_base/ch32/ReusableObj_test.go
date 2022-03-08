package ch32

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewObjPool(t *testing.T) {
	type args struct {
		numOfObj int
	}
	var tests []struct {
		name string
		args args
		want *ObjPool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewObjPool(tt.args.numOfObj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewObjPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjPool_GetObj(t *testing.T) {
	type fields struct {
		bufChan chan *ReusableObj
	}
	type args struct {
		timeout time.Duration
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		want    *ReusableObj
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ObjPool{
				bufChan: tt.fields.bufChan,
			}
			got, err := p.GetObj(tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetObj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjPool_ReleaseObj(t *testing.T) {
	type fields struct {
		bufChan chan *ReusableObj
	}
	type args struct {
		obj *ReusableObj
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ObjPool{
				bufChan: tt.fields.bufChan,
			}
			if err := p.addObj(tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("addObj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// 向里面插入
	//if err := pool.addObj(&ReusableObj{}); err != nil {
	//	t.Error(err)
	//}
	for i := 0; i < 13; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T \n", v)
			// 这里取一个往里面设置一个
			//if err := pool.addObj(v); err != nil {
			//	t.Error(err)
			//}
		}
	}
	fmt.Println("Done")
}
