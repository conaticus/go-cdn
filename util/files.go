package util

import "strings"

// Get file extension from a file
func GetFileExtension(filename string) string {
	split := strings.Split(filename, ".")
	return split[len(split) - 1]
}