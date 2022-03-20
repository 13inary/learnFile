package remove

import "os"

// package remove

// RemoveOne remove one file or directory
func RemoveOne(name string) error {
	err := os.Remove(name)
	if err != nil {
		return err
	}
	return nil
}
