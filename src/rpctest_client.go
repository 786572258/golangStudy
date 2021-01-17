package main

import (
	"fmt"
	"log"
	"net/rpc"
)


// 算数运算请求结构体
type ArithRequest4 struct {
	A int
	B int
}


func main8() {
	rp, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	res := 0
	err2 := rp.Call("Arith.Echo2", ArithRequest{333,444}, &res)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("结果", res)
}