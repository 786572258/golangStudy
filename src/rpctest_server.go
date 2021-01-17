package main

import (
    "fmt"
    "log"
    "net/http"
    "net/rpc"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
    A int
    B int
}



// func (变量 *结构体名) 方法名(变量 类型) 返回类型{}
func (this *Arith) Echo2(req ArithRequest, res *int) error{
    fmt.Println(req.A, req.B)
    *res = req.A + req.B
    return nil
}


func main11() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    err := http.ListenAndServe("127.0.0.1:8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}