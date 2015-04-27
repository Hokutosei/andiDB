package main

import (
	"andiDB/command"
	"fmt"
	"net/http"
    "encoding/json"
)

// Input entrypoint for HTTP client
func Input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled!")
	var c command.Command
	_ = c

    ReadPost(r)
}

// ReadPost extract post request from http client
func ReadPost(r *http.Request) {
    // fmt.Println(r)
    // query := r.URL.Query()
    // fmt.Println(query)


    decoder := json.NewDecoder(r.Body)

    var data interface{}
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(data)
}
