package write

import (
	"io"
	"os"
)

// package write

// WriteString
func WriteString(f *os.File, content string) (int, error) {
	byteCount, err := io.WriteString(f, content)
	if err != nil {
		return 0, err
	}
	return byteCount, nil
}
