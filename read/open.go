package read

import (
	"io/ioutil"
	"os"
)

// package read

// ReadOpen
func ReadOpen(f *os.File) (*[]byte, error) {
	date, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return &date, nil
}
