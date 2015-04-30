package command

import "andiDB/storage"

// LpushCmd push item to left of slice
func LpushCmd(key string, values []interface{}) []string {
	var s []string
	s = storage.SliceDb[key]

	if len(s) == 0 {
		storage.SliceDb[key] = []string{}
		s = storage.SliceDb[key]
	}

	storage.SliceDb[key] = append(s, values[0].(string))
	return storage.SliceDb[key]
}
