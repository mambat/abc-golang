package main

import (
    "net/http"
    "fmt"
    "log"
)

// 服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求
func main()  {
    // HandleFunc vs HandlerFunc
    http.HandleFunc("/", handler) // each request calls handler
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
