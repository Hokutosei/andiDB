package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "testing"
)

// post to server
func Post(cmd string, key string, values []string) {
	url := "http://localhost:8000/post"

	data := map[string]interface{}{
		"cmd":    cmd,
		"key":    key,
		"values": values,
	}
	mJSON, _ := json.Marshal(data)
	content := bytes.NewReader(mJSON)

	req, _ := http.NewRequest("POST", url, content)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "andiDB")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// func BenchmarkPost(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fmt.Sprintf("h: %d", i)
//
// 		Post("set", "server1", "10.0.1.1")
// 	}
// }

func main() {
	argsWithProg := os.Args

	command := argsWithProg[1]
	key := argsWithProg[2]

	// values holder
	values := []string{}
	for i := 3; i < len(argsWithProg); i++ {
		// fmt.Println(argsWithProg[i])
		// fmt.Println("args with prog debug")
		values = append(values, argsWithProg[i])
	}

	Post(command, key, values)
}
