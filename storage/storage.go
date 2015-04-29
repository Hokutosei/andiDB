package storage

// global db handler/ storage
var (
	StringDb = map[string]string{}

	SliceDb = map[string][]interface{}{}
)

type ListStruct map[string]interface{}

// AllDb return all db in slice
func AllDb() []interface{} {
	dbList := []ListStruct{StringDb, SliceDb}

	return dbList
}

func extractKeys() {
	// strDb :=
}
