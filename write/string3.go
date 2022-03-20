package write

import (
	"bufio"
	"os"
)

// package write

// WriteString
func WriteString3(f *os.File, content string) (int, error) {
	w := bufio.NewWriter(f)
	defer w.Flush()

	byteCount, err := w.WriteString(content)
	if err != nil {
		return 0, err
	}
	return byteCount, nil
}
