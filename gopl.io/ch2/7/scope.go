package main

import (
    "os"
    "fmt"
    "log"
)

var cwd string

func init() {
    // 这里使用的 短变量声明，导致内部的 cwd 覆盖了包级别变量 cwd，导致包级别变量 cwd未被初始化
    cwd, err := os.Getwd() // NOTE: wrong!

    // 通过将单独声明 err 和 正常的变量声明，可以避免这个隐藏的 Bug
    //var err error
    //cwd, err = os.Getwd()

    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}



func main() {
    // if 初始化部分 创建了一个 隐式的词法域，创建的变量可以在后续的 else 语句块中访问
    if f, err := os.Open("fname"); err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
    } else {
        // f and err are visible here too
        f.Close()
    }
}
