package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result

}

func test(a string, b int) {
	fmt.Println(a, b)
}

func testMultiReturn(x int , y string) (int, string) {
	return x, y
}

func swap(a *int) {
	*a = 5

	fmt.Println(a)
}

func main15() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!"+"weirui")
   var a, b  string = "1", "2"
   fmt.Println(a, b)

   //var c = "weirui"
   d := "aa" // var d string = "aa"
   fmt.Println(d)

   var (
   		e int
   		f bool
   )

   fmt.Println(e,f)

   g, h := 1, "hello"
   fmt.Println(g, h)

   const LENGTH int = 10
   fmt.Println(LENGTH)

   const (
   		A = 1
   		B string = "bbbb"
   )

   fmt.Println(A, B)
   var i = false


   fmt.Println(i == false)


   for true {
   	fmt.Println("这是无限循环")
   	break
   }

   result := max(5,7);
   fmt.Println("最大值：%d",  result)

   test("ddd", 3)
   fmt.Println( testMultiReturn(555, "weirui"))

   var aa int = 6
   swap(&aa)
   fmt.Println(aa)

   var c1 Circle
   c1.redius = 10.00
   fmt.Println("圆的面积 = ", c1.getArea(3))

   //var arr[3] int
   //arr[0] = 3

   //var arr = [3]int{0,1,2}
   //var arr = [2][3]int{{0,1,2},{3,4,5}}
   var arr = [...]float32{3.2,5.6,4.4}
   //var 变量 = [长度]类型{值,值}
   fmt.Println(arr)
   testPtrArr()


	var arr2 [1]int
	arr2[0] = 1
	//arr2[1] = 2
	fmt.Println(arr2)

	sliceTest()
   sliceTest2()
   rangeTest()

   mapTest()

   result2 := recursion(3)
   fmt.Println(result2)

   var sum uint8 = 255
   fmt.Println(sum)
   fmt.Println(float32(255))

   // 接口的调用
   var phone Phone
   phone = new(NokiaPhone)
   phone.call()

   fmt.Println("结果:", deferTest())
   fmt.Println(ii)

   n,err := errorTest(1)

   fmt.Println(n, err)

}

func errorTest(param int) (n int, err error) {
   defer func() {//必须要先声明defer，否则不能捕获到panic异常
      fmt.Println("defer start")
      if err2 := recover(); err2 != nil {
         fmt.Println(err2) //这里的err其实就是panic传入的内容，"bug"
      }
      fmt.Println("defer end")
   }()

   //var a int = 5/0
   fmt.Println("errorTest 。。。")
   panic("bug") 
   return n,err
}

var ii = 10
func deferTest() int {
   // defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
   // panic在延迟函数之前延迟函数是执行不了的，毕竟抛出异常。
   // panic("panic") 

   fmt.Println("deferTest ii = ", ii)

   var i int = 1
   defer func(i *int) {
      ii++
         fmt.Println("in deferTest ii = ", ii)
      *i = 2
      fmt.Println("in deferTest fun")
         fmt.Println("结果对对对", *i)

   }(&i)


   // 多个defer会逆序执行
   defer func() {
      fmt.Println("in deferTest fun2")
   }()

   fmt.Println("deferTest")

   return ii
}

func recursion(n int) int{
   fmt.Println(n)
    // 3
// 1*2*3*4
   if (n < 2) {
      return 1
   }
   return recursion(n - 1)*n                       
}

func mapTest() {
   // var map1 map[string]string
   var map1 = make(map[string]string)

   map1["aa"] = "11"

   fmt.Println(map1 == nil)

   fmt.Println(map1)
   // a,b := map1["aa"]
   fmt.Println(map1["aa"])

   // 声明并初始化
   map2 := map[string]int{"a":5, "b":6, "d":3}
   fmt.Println(map2)

   i, exist := map2["b"]
   fmt.Println(i, exist)

   for k, v := range map2 {
      fmt.Println(k, v)
   }

   changeMap(map2)
   delete(map2, "b")
   fmt.Println(map2)


}

// map的传递是地址传递
func changeMap(m map[string]int) {
   m["a"] = 444
}

func rangeTest() {
   nums := []int{2,3,4}
   sum := 0
   for i, num := range nums {
      fmt.Println("index:", i)
      sum += num
   }
   fmt.Println("sum:", sum)
}

// 切片
/*golang分配内存有一个make函数，该函数第一个参数是类型，第二个参数是分配的空间，
第三个参数是预留分配空间，前两个参数都很好理解，但我对第三个参数却一脸懵逼，
例如a:=make([]int, 5, 10)， len(a)输出结果是5，cap(a)输出结果是10，
然后我对a[4]进行赋值发现是可以得，但对a[5]进行赋值发现报错了，
于是郁闷这个预留分配的空间要怎么使用呢，
于是google了一下发现原来预留的空间需要重新切片才可以使用*/
func sliceTest() {
    var arr = []string{"aaa", "bbb", "ccc", "ddd"}
	//var arr = make([]int, 5, 10) 
	//arr[2] = 1
	fmt.Println(arr)
	fmt.Printf("length=%d, capacity=%d\n", len(arr), cap(arr))
   // 切片是对数组的引用！！
   var arr2 = arr[2:4] // arr[startIndex:endIndex] 将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片


   fmt.Println("切片arr2:", arr2, cap(arr2))
   fmt.Println(&arr[2])
   arr2[0] = "ffffff"
   fmt.Println(arr, arr2)
   //fmt.Println(&arr2[0])

   fmt.Println(arr2)

   fmt.Printf("length=%d, capacity=%d\n", len(arr2), cap(arr2))


}

/**
 a := [3]int{1, 2, 3}
 b := [...]int{1, 2, 3}
 c := []int{1, 2, 3}
 fmt.Println(cap(a), cap(b), cap(c))
 a = append(a, 4)//Error:first argument to append must be slice; have [3]int
 b = append(b, 4)//Errot:first argument to append must be slice; have [3]int
 c = append(c, 4)//正常，说明变量c是Slice类型
 */
func sliceTest2() {
   fmt.Println("切片 apend的案例")
   var numbers []int
   printSlice(numbers)

   /* 允许追加空切片 */
   numbers = append(numbers, 0)
   printSlice(numbers)

   /* 向切片添加一个元素 */
   numbers = append(numbers, 1)
   printSlice(numbers)

   /* 同时添加多个元素 */
   numbers = append(numbers, 2,3,4)
   printSlice(numbers)

   /* 创建切片 numbers1 是之前切片的两倍容量*/
   numbers1 := make([]int, len(numbers), (cap(numbers))*2)

   /* 拷贝 numbers 的内容到 numbers1 */
   copy(numbers1,numbers)
   printSlice(numbers1)  

}

func printSlice(x []int) {
   fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
// 指针数组
func testPtrArr() {
	arr := []int{10, 20, 30}
	var ptr [3]*int
	for i := 0; i < 3; i++ {
		ptr[i] = &arr[i]
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	*ptr[1] = 222
	fmt.Println(*ptr[1], arr[1])


	//book := Book{title:"php从入门到精通"}
	var book = Book{title:"php从入门到精通"}
	printBook(&book)
	fmt.Println(book.title)

}

func printBook(book *Book) {
	book.title = "改了"

}

type Circle struct {
	redius float64
}

/*
func (变量 变量类型) 方法名(变量 变量类型) 返回类型 {

}
*/
func (c Circle) getArea(a int) float64 {
	return c.redius * 3.14 * c.redius
}

type Book struct {
	title string
	author string
	bookid int
}



// 接口案例
type Phone interface {
   call() int
}

type NokiaPhone struct {

}

/* 实现接口方法 
func (struct_name_variable struct_name) method_name1() [return_type] {
  
}
 */
func (nokiaPhone NokiaPhone) call() int {
   fmt.Println("I am Nokia, I can call you!  ")
   return 1
}













