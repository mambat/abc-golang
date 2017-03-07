// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
    // 函数内部的名字则必须先声明之后才能使用
    const freezingF, boilingF = 32.0, 212.0
    fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
    fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

// 包一级的各种类型的声明语句的顺序无关紧要
// 包一级的声明，首字母大写表示可以在包外访问（比如fmt.Printf），首字母小写则表示只能在包内访问
func fToC(f float64) float64 {
    return (f - 32) * 5 / 9
}
