package check

import "os"

// package check

// Check check path and file
func Check(pathFile string) bool {
	if _, err := os.Stat(pathFile); os.IsNotExist(err) { // not exist
		return false
	}
	return true
}
