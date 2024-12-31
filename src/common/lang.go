package common

import (
	"reflect"
	"strconv"
)

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Float32ToString(f float32)string{
	return strconv.FormatFloat(float64(f), 'E', -1, 32)
}
func Float64ToString(f float64)string{
	return strconv.FormatFloat(float64(f), 'E', -1, 64)
}

func IsTypeOf(unkown interface{}, typeName string) bool {
	if _,ok:=reflect.TypeOf(unkown).FieldByName(typeName);ok{
		return ok
	}
	return false
}
