package moretypes

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	wordList := strings.Fields(s)
	for _, value := range wordList {
		result[value]++
	}
	return result
}

func ExerciseMaps() {
	wc.Test(WordCount)
}
