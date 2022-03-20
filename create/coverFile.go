package create

import (
	"os"
)

// package create

// CreateFile cover exist file and use 0666
func CreateFile(pathFile string) (*os.File, error) {
	f, err := os.Create(pathFile) // create new file
	if err != nil {
		return nil, err
	}
	return f, nil
}
