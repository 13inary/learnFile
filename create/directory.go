package create

import (
	"fmt"
	"io/fs"
	"os"
)

// package create

// CreateDir create more directory. if exist, not cover
// perm 0777
func CreateDir(path string, perm fs.FileMode) error {
	//递归创建文件夹
	fmt.Printf("\n========== perm >>>>>>>>>>\nperm = %+v\n========== perm <<<<<<<<<<\n\n", perm)
	err := os.MkdirAll(path, perm)
	if err != nil {
		return err
	}
	return nil
}
