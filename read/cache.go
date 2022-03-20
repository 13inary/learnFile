package read

import "io/ioutil"

// package read

// ReadCache
func ReadCache(pathName string) ([]byte, error) {
	data, err := ioutil.ReadFile(pathName)
	if err != nil {
		return nil, err
	}
	return data, nil
}
