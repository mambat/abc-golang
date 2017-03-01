package main

import (
    "fmt"
    "os"
)

// cd gopl.io/ch1/
// go run exercise2.go a b c
func main() {
    for idx, arg := range os.Args[1:] {
        fmt.Print(idx)
        fmt.Println(" : " + arg)
    }
}
