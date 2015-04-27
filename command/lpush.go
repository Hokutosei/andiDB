package command

import "andiDB/storage"

// LpushCmd push item to left of slice
func LpushCmd(key string, values []interface{}) []interface{} {
	var s []interface{}
	s = storage.SliceDb[key]

	if len(s) == 0 {
		storage.SliceDb[key] = []interface{}{}
		s = storage.SliceDb[key]
	}

	storage.SliceDb[key] = append(s, values[0])
	return storage.SliceDb[key]
}
