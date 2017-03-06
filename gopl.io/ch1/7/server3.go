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
    http.HandleFunc("/", handler3) // each request calls handler

    // /count 请求不会被 handler2 处理
    http.HandleFunc("/count", counter2)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// counter echoes the number of calls so far.
func counter2(w http.ResponseWriter, request *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

// handler echoes the HTTP request.
func handler3(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    // 简单的语句结果 作为if 语句的最前面：简洁，限制变量作用域
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }

    mu.Lock()
    count++
    mu.Unlock()
    fmt.Println(r.URL.Path)
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
