package main

import (
	"fmt"
	"golangStudy/mysqltest"
	_ "golangStudy/mysqltest"
)

func main() {
	fmt.Println(222)
	//grpcExample.Main2()
	//gintest.RunGinTest()
	//mysqltest.Main()
	mysqltest.Gormtest()
}