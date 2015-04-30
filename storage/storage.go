package storage

// global db handler/ storage
var (
	StringDb = map[string]string{}

	SliceDb = map[string][]string{}
)

// AllDb return all db in slice
func AllDb() []interface{} {
	dbList := []interface{}{StringDb, SliceDb}

	return dbList
}
