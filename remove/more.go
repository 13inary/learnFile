package remove

import "os"

// package remove

// RemoveMore remove all in directory or remove one file
func RemoveMore(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
