package write

import (
	"io/fs"
	"io/ioutil"
)

// package write

// WriteCover cover file
func WriteCover(pathName string, datas []byte, perm fs.FileMode) error {
	err := ioutil.WriteFile(pathName, datas, perm)
	if err != nil {
		return err
	}
	return nil
}
