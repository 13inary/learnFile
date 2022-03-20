package copy

import (
	"io"
	"os"
)

// package copy

// CopyXxxder copy reader to writer
func CopyXxxder(dst *os.File, src *os.File) (int64, error) {
	n, err := io.Copy(dst, src)
	if err != nil {
		return 0, err
	}
	return n, nil
}
