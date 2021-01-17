package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

//主协程比较特殊，如果主协程执行结束后整个程序就要退出，所以 printHello 协程得不到机会去执行下面的输出的语句了，所以以上的程序的数据结果如下：
func test14() {
	time.Sleep(99 * time.Millisecond)

	fmt.Println("hello")

}
func main7() {
	fmt.Println("start")
	go test14()

	fmt.Println(time.Millisecond)
	fmt.Println("end")
		time.Sleep(100*time.Millisecond)



}