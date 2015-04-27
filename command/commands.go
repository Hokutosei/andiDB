package command

import (
	"andiDB/storage"
	"fmt"
)

// Command Database command interface
type Command interface {
	Cmd()
}

const (
	get = "get"
	set = "set"
)

// Cmd command selector
func Cmd(cmd string, key string, values []interface{}) {
	switch cmd {
	case get:
		GetCmd(key)

	case set:
		if len(values) > 1 {
			fmt.Println("must have only one value")
			return
		}

		SetCmd(key, values[0].(string))

	}
}

// GetCmd get command implementation
// @param key string
// @param value string
func GetCmd(key string) {
	result := storage.StringDb[key]
	fmt.Println(result)
}

// SetCmd set command implementation
// @param key string
// @param value string
func SetCmd(key string, value string) {
	storage.StringDb[key] = value
	fmt.Println(0)
}
