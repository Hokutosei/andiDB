package command

import (
	"andiDB/storage"
	"strconv"
)

// LrangeCmd get items in slice by index
func LrangeCmd(key string, values []interface{}) []string {
	// empty slice
	var x []string

	if len(values) < 2 {
		return x
	}

	s, _ := strconv.Atoi(values[0].(string))
	e, _ := strconv.Atoi(values[1].(string))

	item := storage.SliceDb[key]

	// test some condition if
	// user wants whole range
	// or invalid input
	switch {
	case s == 0 && e == -1:
		return item
	case s != 0 && e == -1:
		return x
	case s == -1:
		return x
	case e == 0:
		return x
	case len(item) == 0:
		return x
	}

	i := item[s:e]
	return i
}
