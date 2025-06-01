package kata

import (
	"sort"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	wordSlice := strings.Split(sentence, " ")

	sort.Slice(wordSlice, func(i, j int) bool {
		var num1, num2 int
		w1 := wordSlice[i]
		w2 := wordSlice[j]
		for _, r1 := range w1 {
			if unicode.IsDigit(r1) {
				num1 = int(r1) - 48
			}
		}

		for _, r2 := range w2 {
			if unicode.IsDigit(r2) {
				num2 = int(r2) - 48
			}
		}

		return num1 < num2
	})

	result := strings.Join(wordSlice, " ")

	return result
}
