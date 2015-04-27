package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// post to server
func post(cmd string, key string, values []string) {
	v := url.Values{}
	v.Set("command", cmd)
	v.Add("key", key)

	for i, item := range values {
		hashKey := fmt.Sprintf("%s", string(i))
		v.Add(hashKey, string(item))
	}

	fmt.Println(v)

	// resp, err := http.PostForm("http://localhost:8000/post", url.Values{"command": {cmd}, "key": {key}, "values": {values}})
	resp, err := http.PostForm("http://localhost:8000/post", v)
	if err != nil {
		fmt.Println("err post: ", err)
		return
	}

	fmt.Println(resp)
}

func main() {
	fmt.Println("post")

	argsWithProg := os.Args

	command := argsWithProg[1]
	key := argsWithProg[2]
	fmt.Println(command)

	// values holder
	values := []string{}
	for i := 3; i < len(argsWithProg); i++ {
		values = append(values, argsWithProg[i])
	}

	fmt.Println(values)
	post(command, key, values)
}
