package command

import (
	"fmt"
	"time"
)

// Command Database command interface
type Command interface {
	Cmd()
}

const (
	get    = "get"
	set    = "set"
	lpush  = "lpush"
	lrange = "lrange"

	keys = "keys"
)

// ResponseData data struct to return to user
type ResponseData struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

// Cmd command selector
// mechanism for distinguishing command/task
func Cmd(cmd string, key string, values []interface{}, resultChan chan ResponseData) {
	start := time.Now()
	switch cmd {

	// get command condition
	case get:
		// get()
		resultChan <- ResponseData{200, GetCmd(key)}

	// set command condition
	case set:
		// validate if values length has more than 1
		if len(values) > 1 {
			fmt.Println("must have only one value")

			resultChan <- ResponseData{500, ""}
			return
		}

		// set()
		d := ResponseData{200, SetCmd(key, values[0].(string))}
		resultChan <- d

	case lpush:
		// lpush()
		resultChan <- ResponseData{200, LpushCmd(key, values)}

	case lrange:

		// lrange()
		resultChan <- ResponseData{200, LrangeCmd(key, values)}

	case keys:

		// keys()
		resultChan <- ResponseData{200, KeysCmd(key, values)}
	}

	fmt.Println("Cmd took: ", time.Since(start))

}
