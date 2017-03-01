// Echo1 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

// for 循环的形式一：
// for initialization; condition; post {
//     // zero or more statements
// }
func main() {
    // 被隐式地赋予其类型的零值（zero value）
    var s, sep string

    // os.Args[0] 是命令本身的名字（命令行执行 go run echo1.go 查看结果）
    fmt.Println(os.Args[0])

    // 符号 := 是短变量声明（short variable declaration）的一部分
    // i++ 是语句，而不像C系的其它语言那样是表达式（j = i++ 是非法的）；而且 ++ 和 -- 都只能放在变量名后面，因此--i也非法。
    // initalization （这里的i := 1） 如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}
