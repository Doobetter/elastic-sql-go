package common

import (
	"fmt"
	"testing"
)

type A1 struct {
	Name string
}
type S1 struct {
	A1
	Age int
}
type SS1 struct {
	S1
	Gender int
}

func TestIsTypeOf(t *testing.T) {
	//fmt.Println()
	ss1:=SS1{
		S1:     S1{
			A1:  A1{Name:"aaa"},
			Age: 2,
		},
		Gender: 3,
	}
	un:= interface{}(ss1)
	fmt.Println(IsTypeOf(un,"S1"))
}
