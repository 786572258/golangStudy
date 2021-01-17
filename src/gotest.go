// 错误处理
package main
import(
    "fmt"
    "errors"
)

var errDivByZero error = errors.New("division by zero")

func div(dividend, divisor int) (int, error){
    if divisor == 0{
        return 0, errDivByZero
    }
    return dividend / divisor, nil
}

func main9(){
    fmt.Println(div(1, 0))
    fmt.Println(div(6, 2))
}