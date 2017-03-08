package tempconv

import "fmt"

func CToF(c Celsius) Fahrenheit {
    // 类型转换操作
    return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
    // 类型转换操作
    return Celsius((f - 32) * 5 / 9)
}

func main()  {
    // 包级别的声明，例如在一个文件声明的类型和常量，在同一个包的其他源文件也是可以直接访问的，就好像所有代码都在一个文件一样
    fmt.Printf("Brrrr! %v\n", AbsoluteZeroC)
}