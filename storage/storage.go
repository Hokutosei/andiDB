package storage

// global db handler/ storage
var (
	StringDb = map[string]string{}

	SliceDb = map[string][]interface{}{}
)

// AllDb return all db in slice
func AllDb() []interface{} {
	db := []interface{}{StringDb, SliceDb}

	return db
}
