package main

import (
	"fmt"
	"net/http"
)

const (
	serverPort = ":8000"
)

// main program entrypoint
func main() {

	go AsyncDumpDb()

	// print server data status/info
	go func() {
		fmt.Println("server listening to: ", serverPort)
	}()

	http.HandleFunc("/post", Input)
	http.ListenAndServe(serverPort, nil)
}
