package main

import (
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"
)

// post to server
func post(cmd string, key string, values []string) {
	url := "http://localhost:8000/post"

	data := map[string]interface{}{
		"cmd": cmd,
		"key": key,
		"values": values,
	}
	mJSON, _ := json.Marshal(data)
	content := bytes.NewReader(mJSON)

	req, _ := http.NewRequest("POST",url, content)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "andiDB")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response: ", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body: ", string(body))
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
