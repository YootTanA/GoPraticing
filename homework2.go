package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	wordList := strings.Split(s, " ")
	wordMapper := make(map[string]int)

	for i := range wordList {

		word := strings.Trim(wordList[i], ",")

		if val, ok := wordMapper[word]; ok {

			wordMapper[word] = val + 1

		} else {
			wordMapper[word] = 1
		}
	}

	return wordMapper
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quack like a duck, then it probably is a duck"
	w := WordCount(s)
	fmt.Println(w)
}
