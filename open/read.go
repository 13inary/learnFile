package open

import "os"

// package open

// OpenRead
func OpenRead(pathName string) (*os.File, error) {
	f, err := os.Open(pathName)
	if err != nil {
		return nil, err
	}
	return f, nil
}
