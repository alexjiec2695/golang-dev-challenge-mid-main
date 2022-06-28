package tfidf

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"
)

func Run(a string) (float64, error) {

	var files []fs.DirEntry
	var err error
	var wordInDocument, large, count float64
	var word string

	files, err = os.ReadDir(a)

	if err != nil {
		return 0, fmt.Errorf("error file not found: %s", err)
	}

	documents := float64(len(files))

	for _, file := range files {
		var bytes []byte

		bytes, err = os.ReadFile(fmt.Sprintf("%s/%s", a, file.Name()))

		if err != nil {
			return 0, fmt.Errorf("error in read file %s", err)
		}

		content := string(bytes)
		count += float64(strings.Count(strings.ToLower(content), strings.ToLower(word)))
		large += float64(len(content))

		if strings.Count(strings.ToLower(content), strings.ToLower(word)) > 0 {
			wordInDocument += 1
		}
	}

	tf := count / large

	idf := math.Log(documents / (wordInDocument + 1))

	tfIdf := tf * idf

	return tfIdf, nil
}
