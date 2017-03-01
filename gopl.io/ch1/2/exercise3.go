package main

import (
    "fmt"
    "os"
    "time"
    "strings"
)

// cd gopl.io/ch1/
// go run exercise2.go a b c
func main() {
    var start = time.Now().UnixNano()
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
    durationFor := time.Now().UnixNano() - start

    start = time.Now().UnixNano()
    fmt.Println(strings.Join(os.Args[0:], " "))
    durationJoin := time.Now().UnixNano() - start

    fmt.Printf("durationFor: %d nanoseconds\n", durationFor)
    fmt.Printf("durationJoin: %d nano", durationJoin)
}
