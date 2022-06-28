package main

import (
	"fmt"
	"golang-dev-challenge-mid/pkg/tfidf"
)

func main() {
	run, err := tfidf.Run("../../data")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR: %s", err))
	}
	fmt.Println("TFIDF: ", run)

}
