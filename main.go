package main

import (
	"learnFile/create"
	"os"
)

// package main

// main
func main() {
	f, err := create.CreateFile("./new.txt")
	if err != nil {
		return
	}
	defer f.Close()

	os.RemoveAll("./test/aa.txt")
}
