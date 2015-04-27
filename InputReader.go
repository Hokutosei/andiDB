package main

import (
	"andiDB/command"
	"encoding/json"
	"fmt"
	"net/http"
)

// InputPost struct handler for user input
type InputPost struct {
	Command string        `json:"cmd"`
	Key     string        `json:"key"`
	Values  []interface{} `json:"values"`
}

// Input entrypoint for HTTP client
func Input(w http.ResponseWriter, r *http.Request) {
	inputPost, err := ReadPost(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	command.Cmd(inputPost.Command, inputPost.Key, inputPost.Values)
}

// ReadPost extract post request from http client
func ReadPost(r *http.Request) (InputPost, error) {
	decoder := json.NewDecoder(r.Body)

	var data InputPost
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	fmt.Println(data)
	return data, nil
}
