package main

import (
	"fmt"
	"net/http"

	"andiDB/command"
)

// Input entrypoint for HTTP client
func Input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handled!")
	var c command.Command
	_ = c
}

// main program entrypoint
func main() {
	fmt.Println("started!")

	http.HandleFunc("/post", Input)
	http.ListenAndServe(":8000", nil)
}
