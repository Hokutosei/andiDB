package command

// SetCmd set command implementation
import "andiDB/storage"

// @param key string
// @param value string
func SetCmd(key string, value string) string {
	storage.StringDb[key] = value
	return "ok"
}
