package main

import "fmt"

func test2() {
	fmt.Println(3434)
}

func main19() {
	slice1 := make([]float32, 2, 2) // 长度为0的切片
	//slice1[0] = 33
	//slice1[1] = 44
	slice1 = append(slice1, 1)

	fmt.Println(slice1, cap(slice1))
}
