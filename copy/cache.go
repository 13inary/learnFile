package copy

import (
	"io/fs"
	"io/ioutil"
)

// package copy

// CopyCache copy []byte to pathName
func CopyCache(pathName string, src []byte, perm fs.FileMode) error {
	err := ioutil.WriteFile(pathName, src, perm)
	if err != nil {
		return err
	}
	return nil
}
