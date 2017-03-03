package main

import (
    "os"
    "net/http"
    "fmt"
    "io"
    "strings"
)

func main() {
    urlHttpPrefix := "http://"
    for _, url := range os.Args[1:] {
        // exercise 1.8
        if !strings.HasPrefix(url, urlHttpPrefix) {
            url = urlHttpPrefix + url
        }

        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            os.Exit(1)
        }

        //b, err := ioutil.ReadAll(resp.Body)

        // exercise 1.7
        io.Copy(os.Stdout, resp.Body)

        // exercise 1.9
        fmt.Fprintf(os.Stdout, "http status code: %d", resp.StatusCode)

        resp.Body.Close()

        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }

        //fmt.Printf("%s", b)
    }
}
