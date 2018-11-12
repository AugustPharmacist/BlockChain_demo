package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := int16(0x1234)
	b := int8(s)
	//0x1234
	//低------》高
	//12 34 -》 大端 -》 高尾端
	//34 12 -》 小段 -》 低尾端

	fmt.Println("int16字节大小为", unsafe.Sizeof(s)) //结果为2
	if 0x34 == b {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}
}
