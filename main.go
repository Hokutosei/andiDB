package main

import (
	"fmt"
	"net/http"
)

// main program entrypoint
func main() {
	fmt.Println("started!")

	http.HandleFunc("/post", Input)
	http.ListenAndServe(":8000", nil)
}
