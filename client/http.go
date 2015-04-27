package main

import (
	"fmt"
	"os"
)

func post() {

}

func main() {
	fmt.Println("post")

	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	command := argsWithProg[1]
	fmt.Println(command)

	// values holder
	values := []string{}
	for i := 2; i < len(argsWithProg); i++ {
		values = append(values, argsWithProg[i])
	}

	fmt.Println(values)
}
