package main

import (
    "time"
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
)

// 当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，
// 直到另一个goroutine往这个channel里写入、或者接收值，这样两个goroutine才会继续执行channel操作之后的逻辑
func main() {
    start := time.Now()
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }

    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }

    // time.Since
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)

    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    // ioutil.Discard 只关心字节数，不关心内容
    //nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    //resp.Body.Close() // don't leak resources
    //if err != nil {
    //    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    //    return
    //}
    //
    //secs := time.Since(start).Seconds()
    //ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

    // exercise 1.10
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    //fmt.Printf("%s", b)

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(b), url)
}
