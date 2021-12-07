package main

import (
	"flag"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func delete(filename string) {
	defer wg.Done()

	err := os.RemoveAll(filename)

	if err != nil {
		panic(err)
	}
}

func main() {
	globPointer := flag.String("glob", "", "Glob pattern to match")
	flag.Parse()

	filesFoldersList, err := filepath.Glob(*globPointer)

	if err != nil {
		panic(err)
	}

	finalFilesFoldersList := make([]string, 0)
	for _, filename := range filesFoldersList {
		stat, err := os.Stat(filename)

		if err != nil {
			panic(err)
		}

		if stat.IsDir() {
			list, err := filepath.Glob(filename + "/*")

			if err != nil {
				panic(err)
			}

			finalFilesFoldersList = append(finalFilesFoldersList, list...)
		}

		finalFilesFoldersList = append(finalFilesFoldersList, filename)
	}

	for _, filename := range finalFilesFoldersList {
		wg.Add(1)

		go delete(filename)
	}

	wg.Wait()
}
