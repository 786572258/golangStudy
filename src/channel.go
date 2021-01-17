package main
import "fmt"
import "time"

// 接收(<-channel)或者发送(channel <- xxx)的

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	time.Sleep(1000 * time.Millisecond)
	// 向channel 发送数据
	c <- sum
}
func main18() {
	// ch := make(chan int, 1) // 创建无缓冲channel
	// go func() {

	//   fmt.Println("time sleep 2 second...")
	//   time.Sleep(2 * time.Second)

	//   <-ch

	// }()

	// fmt.Println("即将阻塞...")

	// ch <-1  // 协程将会阻塞，等待数据被读取

	// fmt.Println("ch 数据被消费，主协程退出")


	ch := make(chan int) // 创建有缓冲channel

	go func() {

	  //fmt.Println("time sleep 2 second...")

	  time.Sleep(2 * time.Second)
	  for t := range ch {
	  	fmt.Println("ch值：", t)
	  }
	  // <-ch

	}()

	ch <-1  // 协程不会阻塞，等待数据被读取

	fmt.Println("第二次发送数据到channel， 即将阻塞")
	//time.Sleep(1 * time.Second)

	ch <-2  // 第二次发送数据到channel, 在数据没有被读取之前，因为缓冲区满了， 所以会阻塞主协程。
	ch <-3
	ch <-4

	fmt.Println("ch 数据被消费，主协程退出")
	time.Sleep(10 * time.Second)


}

func worker(done chan bool) {
    time.Sleep(time.Second)
    // 通知任务已完成
    done <- true
}
// channel可以用在goroutine之间的同步。

func main4() {
    done := make(chan bool, 1)
    go worker(done)
    // 等待任务完成
    <-done
}

func main3() {
    go func() {
        time.Sleep(1 * time.Hour)
    }()
    c := make(chan int)
    go func() {
        for i := 0; i < 10; i = i + 1 {
            c <- i
        }
        close(c)
    }()
    for i := range c {
        fmt.Println(i)
    }
    fmt.Println("Finished")
}

func main2() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)   // 7 2 8
	go sum(s[len(s)/2:], c)	  // -9, 4, 0
	
	// 从channel 取数。x, y := <-c, <-c这句会一直等待计算结果发送到channel中。
	x, y := <-c, <-c

	//x, ok := <-c
	fmt.Println(x,y)


}