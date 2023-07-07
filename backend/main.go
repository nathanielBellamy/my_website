package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Printf("Starting server on 8080")
    fs := http.FileServer(http.Dir("./../frontend/dist/"))

    http.Handle("/", fs)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
