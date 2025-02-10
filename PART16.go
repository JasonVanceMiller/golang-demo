//BASIC SERVER

package main

import (
    "fmt"
    "net/http"
)

func  main() {
    fmt.Println("Running")
    http.HandleFunc("/", Hello)
    http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World\n")
}
