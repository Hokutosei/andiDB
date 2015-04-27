package command

import "fmt"

// Command Database command interface
type Command interface {
	Cmd()
}

const (
	get    = "get"
	set    = "set"
	lpush  = "lpush"
	lrange = "lrange"
)

// ResponseData data struct to return to user
type ResponseData struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

// Cmd command selector
// mechanism for distinguishing command/task
func Cmd(cmd string, key string, values []interface{}, resultChan chan ResponseData) {
	switch cmd {

	// get command condition
	case get:
		resp := GetCmd(key)
		resultChan <- ResponseData{200, resp}

	// set command condition
	case set:
		// validate if values length has more than 1
		if len(values) > 1 {
			fmt.Println("must have only one value")
			d := ResponseData{500, ""}
			resultChan <- d
			return
		}
		resp := SetCmd(key, values[0].(string))
		d := ResponseData{200, resp}
		resultChan <- d

	case lpush:
		// lpush()

	}
}
