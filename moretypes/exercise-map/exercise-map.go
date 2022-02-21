package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		wordCounts[word] += 1
	}
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
