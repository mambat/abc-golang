// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

// for 循环的形式2：
// 在某种数据类型的区间（range）上遍历
func main() {
    // 使用 短变量声明 来声明并初始化 s 和 sep
    s, sep := "", ""

    // 以下声明是等价的：
    // s := "" // 最简洁，但只能用在函数内部，而不能用于包变量
    // var s string // 依赖于字符串的默认初始化零值机制，被初始化为""
    // var s = "" // 用得很少，除非同时声明多个变量（var a, b = "", ""）
    // var s string = "" // 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了

    // range产生一对值：索引以及在该索引处的元素值，但 range 语法要求，要处理元素, 必须处理索引。
    // 但 Go 语言不允许使用无用的局部变量（local variables），会导致编译错误。
    // 这种情况的解决方法是用空标识符（blank identifier），即 _（也就是下划线）。
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
