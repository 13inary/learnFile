package open

import (
	"io/fs"
	"os"
)

// package check

// CheckOpenFile open a file by flag and perm h. If file not exist, create file
// pathFile path + file name
// flag : user | to divide
// 		os.O_RDONLY  read only
// 		os.O_WRONLY  write only
// 		os.O_RDWR    read and write
// 		os.O_APPEND  write by add
// 		os.O_CREATE  create when file not exist
// 		os.O_EXCL    和O_CREATE配合使用，文件必须不存在
// 		os.O_SYNC    打开文件用于同步I/O
// 		os.O_TRUNC   如果可能，打开时清空文件
// perm 0xxx
func CheckOpenFile(pathFile string, flag int, perm fs.FileMode) (*os.File, error) {
	f, err := os.OpenFile(pathFile, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
