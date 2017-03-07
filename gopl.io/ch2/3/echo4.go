package main

import (
    "flag"
    "fmt"
    "strings"
)

// 指向对应命令行标志参数变量的指针
// -n 用于忽略行尾的换行符
// -s sep 用于指定分隔字符（默认是空格）
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
    // 在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）
    flag.Parse()

    // flag.Args() 返回 非标志参数的普通命令行参数
    fmt.Print(strings.Join(flag.Args(), *sep))

    if !*n {
        fmt.Println()
    }

    // test:
    // ./echo4 a bc def
    // ./echo4 -s / a bc def
    // ./echo4 -n a bc def
    // ./echo4 -help
}
