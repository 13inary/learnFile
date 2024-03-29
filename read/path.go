package read

import "io/ioutil"

// package read

// ReadCache
func ReadPath(pathName string) ([]byte, error) {
	datas, err := ioutil.ReadFile(pathName)
	if err != nil {
		return nil, err
	}
	return datas, nil
}
