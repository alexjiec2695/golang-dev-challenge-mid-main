package tfidf

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"
	"sync"
)

type calculate struct {
	wordInDocument, large, count float64
}

func Run(path string) (float64, error) {
	var files []fs.DirEntry
	var err error
	var wordInDocument, large, count float64

	files, err = os.ReadDir(path)

	if err != nil {
		return 0, fmt.Errorf("error file not found: %s", err)
	}

	documents := float64(len(files))

	c := make(chan calculate, len(files))
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go Calculate(path, file, c, &wg)
	}

	for i := 0; i < len(files); i++ {
		v := <-c
		count += v.count
		large += v.large
		wordInDocument += v.wordInDocument
	}

	tf := count / large

	idf := math.Log(documents / (wordInDocument + 1))

	tfIdf := tf * idf

	return tfIdf, nil
}

func Calculate(path string, file fs.DirEntry, c chan calculate, wg *sync.WaitGroup) {
	defer wg.Done()

	var bytes []byte
	var word string
	var ca calculate

	bytes, err := os.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()))

	if err != nil {
		fmt.Errorf("error in read file %s", err)
	}

	content := string(bytes)
	ca.count = float64(strings.Count(strings.ToLower(content), strings.ToLower(word)))
	ca.large = float64(len(content))

	if strings.Count(strings.ToLower(content), strings.ToLower(word)) > 0 {
		ca.wordInDocument = 1
	}

	c <- ca
}
