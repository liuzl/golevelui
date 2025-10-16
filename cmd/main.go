package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/liuzl/golevelui"
	"github.com/syndtr/goleveldb/leveldb"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsNotExist(err)
	}
	return true
}

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var paths arrayFlags

var files []*leveldb.DB

func main() {
	flag.Var(&paths, "path", "leveldb dir path")
	flag.Parse()

	if len(paths) == 0 {
		flag.Usage()
		return
	}

	defer func() {
		for _, file := range files {
			if err := file.Close(); err != nil {
				fmt.Printf("failed to close db: %v\n", err)
			}
		}
	}()

	for _, path := range paths {
		if !fileExists(path) {
			fmt.Println("file not found: " + path)
			return
		}
		file, err := leveldb.OpenFile(path, nil)
		if err != nil {
			panic(err)
		}
		// Add initial data for testing
		_ = file.Put([]byte("initial-key"), []byte("initial-value"), nil)

		files = append(files, file)

		golevelui.GetLevelAdmin().Register(file, path)
	}

	err := golevelui.GetLevelAdmin().Start()
	if err != nil {
		panic(err)
	}

	<-make(chan any)
}
