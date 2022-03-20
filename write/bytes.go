package write

import "os"

// package write

// WriteBytes
func WriteBytes(f *os.File, data []byte) (int, error) {
	byteCount, err := f.Write(data)
	if err != nil {
		return 0, err
	}
	return byteCount, nil
}
