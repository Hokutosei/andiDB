package command

import (
	"andiDB/storage"
	"fmt"
)

// KeysCmd dump all database keys in list/ slice
func KeysCmd(key string, values []interface{}) {
	dbList := storage.AllDb()

	fmt.Println(dbList)
	fmt.Println(len(dbList))

	keys := make([]string, 0, len(dbList))
	for i := 0; i < len(dbList); i++ {
		iDb := dbList[i]
		fmt.Println(iDb)
		for k := range iDb {
			keys = append(keys, k)
		}
	}
	fmt.Println(keys)
}
