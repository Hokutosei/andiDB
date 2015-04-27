package command

// GetCmd get command implementation
import "andiDB/storage"

// @param key string
// @param value string
func GetCmd(key string) string {
	result := storage.StringDb[key]
	return result
}
