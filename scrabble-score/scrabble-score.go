// Package scrabble calculates scrabble scores of a word
package scrabble

import "strings"

type LetterScore struct {
	letters string
	score   int
}

var (
	ls1   = LetterScore{"AEIOULNRST", 1}
	ls2   = LetterScore{"DG", 2}
	ls3   = LetterScore{"BCMP", 3}
	ls4   = LetterScore{"FHVWY", 4}
	ls5   = LetterScore{"K", 5}
	ls8   = LetterScore{"JX", 8}
	ls10  = LetterScore{"QZ", 10}
	rules = []LetterScore{ls1, ls2, ls3, ls4, ls5, ls8, ls10}
)

// Score calculates scrabble scores of a word
func Score(in string) int {
	total := 0

	for _, c := range in {
		for _, rule := range rules {
			if strings.Contains(rule.letters, strings.ToUpper(string(c))) {
				total += rule.score
				break
			}
		}
	}

	return total
}
