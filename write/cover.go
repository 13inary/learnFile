package write

import (
	"io/fs"
	"io/ioutil"
)

// package write

// WriteCover cover file
func WriteCover(pathName string, data []byte, perm fs.FileMode) error {
	err := ioutil.WriteFile(pathName, data, perm)
	if err != nil {
		return err
	}
	return nil
}
