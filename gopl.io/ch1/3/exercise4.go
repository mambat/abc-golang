package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines2(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines2(f, counts)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines2(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
        if counts[input.Text()] > 1 {
            fmt.Printf("duplicated line occured in %s\n", f.Name())
        }
    }
    // NOTE: ignoring potential errors from input.Err()
}
