package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola world!")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}