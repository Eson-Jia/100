package main

import (
	"sync/atomic"
	"testing"
)

// Test InterleavePrintOneChan
func TestInterleavePrint(t *testing.T) {
	InterleavePrintOneChan()
}

func Test7(t *testing.T) {
	type Param map[string]interface{}
	type Show struct {
		Param
	}
	func() {
		s := new(Show)
		s.Param["RMB"] = 100
	}()
}

func Test9(t *testing.T) {
	type People struct {
	}
}

func Test12(t *testing.T) {
	var value int32
	SetValue := func(delta int32) {
		for true {
			v := value
			if atomic.CompareAndSwapInt32(&value, v, v+delta) {
				break
			}
		}
	}
	SetValue(12)
}
