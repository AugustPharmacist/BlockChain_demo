package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  uint64
}

func main() {

	//定义Person结构，进行编码后传输
	hasaki := Person{"hasaki", 33}

	var buf bytes.Buffer

	//1.定义编码器
	encoder := gob.NewEncoder(&buf)
	//2.使用编码器编码
	err := encoder.Encode(&hasaki)
	//一定记得校验
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("编码后的数据：%x\n", buf.Bytes())

	//============================= 在对端，进行编码
	//对端解码时使用容器
	var p Person

	//1.定义一个解码器
	decoder := gob.NewDecoder(bytes.NewReader(buf.Bytes()))

	var buffer bytes.Buffer
	buffer.Write(buf.Bytes())
	decoder = gob.NewDecoder(&buffer)

	//2.将传过来的字节流进行解码
	decoder.Decode(&p)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("p : %v\n", p)
}
