package get

import "os"

// package get

// GetPerm
func GetPerm(f *os.File) (os.FileMode, error) {
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return info.Mode(), nil
}
