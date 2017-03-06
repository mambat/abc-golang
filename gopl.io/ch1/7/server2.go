// Server2 is a minimal "echo" and counter server.
package main

import (
    "net/http"
    "fmt"
    "log"
    "sync"
)

// 保证每次修改 count 的最多只能有一个 goroutine
var mu sync.Mutex
var count int

func main() {
    // HandleFunc vs HandlerFunc
    // /favicon.ico 请求 会导致计数异常
    http.HandleFunc("/", handler2) // each request calls handler
    // /count 请求不会被 handler2 处理
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, request *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

// handler echoes the Path component of the requested URL.
func handler2(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()
    fmt.Println(r.URL.Path)
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
