package main

import (
	"fmt"
	"golangStudy/testimport"
)

type Man struct {
	height float64
}

type Wom struct {
	height float64
}


// 实现接口
func (m Man) Eat() string {
	fmt.Println("男人吃东西")
	return "男人吃东西"
}


// 实现接口
func (m Man) Jump() {
	fmt.Println("男人跳")
}


func (w Wom) talk() string {
	fmt.Println("女人说话")
	return "女人说话"
}



type Test interface {
}

type MyString string

type Ts struct {
	Num int64
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main5() {
	fmt.Println("接口示例")
	var p testimport.Person
	//fmt.Println(p.Eat())
	fmt.Printf("type: %T", p)
	man := Man{}
	man.Eat()

	var p2 testimport.Person = Man{}
	fmt.Printf("type p2: %T", p2)
	p2.Eat()

	var wom Wom
	wom.height = 160
	wom.talk()

	var mystring MyString = MyString("sdfsdf")
	fmt.Println(mystring)

	var i interface{} = 1
    fmt.Println("i = ", i)
	fmt.Printf("type i: %T ", i)


    i = "abc"
    fmt.Println("i = ", i)
    fmt.Printf("type i: %T  ", i)

    // 空接口，可以实现任意类型
    countryCapitalMap := map[string]interface{}{"France": 123, "Italy": 456, "Japan": 789, "India": "New delhi"}
	fmt.Println(countryCapitalMap)

}