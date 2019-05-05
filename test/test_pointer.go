package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

type MyInterface interface {
}

type MyStruct struct {
	i1 int64
	i2 int64
	i3 int64
	i4 int64
	i5 int64
}

func Hello1(p MyStruct) {
	fmt.Println("size:", unsafe.Sizeof(p), "; type:", reflect.TypeOf(p), "; value:", reflect.ValueOf(p))
}

func Hello2(p * MyStruct) {
	fmt.Println("size:", unsafe.Sizeof(p), "; type:", reflect.TypeOf(p), "; value:", reflect.ValueOf(p))
}

func Hello3(p MyInterface) {
	fmt.Println("size:", unsafe.Sizeof(p), "; type:", reflect.TypeOf(p), "; value:", reflect.ValueOf(p))
}

//func Hello4(p * MyInterface) {
//   fmt.Println("size:", unsafe.Sizeof(p), "; type:", reflect.TypeOf(p), "; value:", reflect.ValueOf(p))
//}

func main() {
	m := MyStruct { 11,22,33,44,55 }
	Hello1(m)
	Hello2(&m)
	Hello3(m)
	//Hello4((* MyInterface)(&m))
}
