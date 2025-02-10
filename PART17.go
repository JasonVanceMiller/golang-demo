//BASIC+ SERVER

package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Payload struct {
    Name string
}

func  main() {
    fmt.Println("Running")
    http.HandleFunc("/", Hello)
    http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
    //Unwrap request body
    if r.Body == nil {
        fmt.Fprintf(w, "Empty request body")
        return
    }
    defer r.Body.Close()

    //Cast body to struct payload
    decoder := json.NewDecoder(r.Body)
    var payload Payload
    err := decoder.Decode(&payload) 
    if err != nil {
        fmt.Fprintf(w, "Failed to decode body")
        return
    }

    //Respond
    fmt.Fprintf(w, "Hello %s\n", payload.Name)
}
