package command

import "andiDB/storage"

// GetCmd get command implementation
// @param key string
// @param value string
func GetCmd(key string) string {
	result := storage.StringDb[key]
	return result
}
