package read

import "os"

// package read

// ReadBytes
func ReadBytes(f *os.File) (int, *[]byte, error) {
	data := []byte{}
	n, err := f.Read(data)
	if err != nil {
		return 0, nil, err
	}
	return n, &data, nil
}
