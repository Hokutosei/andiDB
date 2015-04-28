package command

import "andiDB/storage"

// SetCmd set command implementation
// @param key string
// @param value string
func SetCmd(key string, value string) string {
	storage.StringDb[key] = value
	return "ok"
}
