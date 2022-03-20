package write

import (
	"os"
)

// package write

// WriteString
func WriteString2(f *os.File, content string) (int, error) {
	byteCount, err := f.WriteString(content)
	if err != nil {
		return 0, err
	}
	return byteCount, nil
}
