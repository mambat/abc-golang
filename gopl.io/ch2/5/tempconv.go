package tempconv

import "fmt"

// 定义类型
// 不可以被相互比较 或 混在一个表达式运算
// 刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

// 类型转换操作 T(x)
// 如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)

// 只有当两个类型的底层基础类型相同时（int-float, string-[]byte 等也是可以的），才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，
// 这些转换只改变类型而不会影响值本身

const (
    AbsoluteZeroC Celsius = -273.15 // 绝对零度
    FreezingC     Celsius = 0       // 结冰点温度
    BoilingC      Celsius = 100     // 沸水温度
)

// 声明 Celsius 类型的一个叫名叫 String 的方法，注意这里的 (c Celsius)
func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
    return fmt.Sprintf("%g°F", f)
}

//func main()  {
//    var c Celsius
//    var f Fahrenheit
//    fmt.Println(c == Celsius(f))
//}
