package main

import (
    "fmt"

    "time"
)

func main13() {
	// testPCB()
	testTimeout()
}

/*
检查channel读写超时，并做超时的处理
 */
func testTimeout() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-time.After(time.Second * 3):
				quit <- true
				fmt.Println("超时，通知主线程退出")
				return
			}
		}
	}()

	//fmt.Println(time.Second * time.Duration(3))
	for i := 0; i < 3; i++ {
		g <- i
	}

	<-quit
	fmt.Println("收到退出通知，主线程退出")
}


/**
生产者消费者问题
 */
func testPCB() {
	fmt.Println("test PCB")

	intchan := make(chan int)
	quitChan := make(chan bool)
	quitChan2 := make(chan bool)

	value := 0

	go func() {
		for i := 0; i < 3; i++ {

			value = value + 1
			intchan <- value

			fmt.Println("write finish, value ", value)

			time.Sleep(time.Second)
		}
		quitChan <- true
	}()
	go func() {
		for {
			select {
			case v := <-intchan:
				fmt.Println("read finish, value ", v)
			case <-quitChan:
				quitChan2 <- true
				return
			}
		}

	}()

	<-quitChan2
	fmt.Println("task is done ")
}