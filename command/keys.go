package command

import (
	"andiDB/storage"
	"fmt"
	"reflect"
)

// MapOfString type for map string use for string extraction
type MapOfString map[string]string

// MapOfSlices type for map of slices use for string extraction
type MapOfSlices map[string][]interface{}

// KeysCmd dump all database keys in list/ slice
func KeysCmd(key string, values []interface{}) []string {
	dbList := storage.AllDb()

	fmt.Println(dbList)
	fmt.Println(len(dbList))

	var keys []string
	for i := 0; i < len(dbList); i++ {
		iDb := dbList[i]

		switch t := iDb.(type) {

		case map[string]string:
			extractKeysOfString(iDb.(map[string]string), &keys)

		case map[string][]interface{}:
			extractKeysFromSliceIntrface(iDb.(map[string][]interface{}), &keys)

		default:
			_ = t
		}

	}
	fmt.Println("keys total--")
	fmt.Println(keys)
	return keys
}

// typeof return string representation of object type
func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// extractKeysOfString use to extract keys from
// a map slice string
func extractKeysOfString(v MapOfString, keys *[]string) {
	for k := range v {
		*keys = append(*keys, k)
	}
}

// extractKeysFromSliceIntrface use to extract
// keys from a map of slices interface
func extractKeysFromSliceIntrface(v MapOfSlices, keys *[]string) {
	for k := range v {
		*keys = append(*keys, k)
	}
}
