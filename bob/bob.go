// Package bob generates teenager's responses
package bob

// import regexp library
import "regexp"

// Hey generates teenager's responses
func Hey(remark string) string {
	switch {
	// check for silence
	case MatchString("^\\s*$", remark):
		return "Fine. Be that way!"
	// check for yelled question
	case MatchString("^([^a-z]*[A-Z][^a-z]*)+\\?$", remark):
		return "Calm down, I know what I'm doing!"
	// check for yelled statement
	case MatchString("^([^a-z]*[A-Z][^a-z]*)+$", remark):
		return "Whoa, chill out!"
	// check for all other questions
	case MatchString("^.*\\?\\s*$", remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

// MatchString is a wrapper for regexp.MatchString so it can be used in switch
func MatchString(pattern, s string) bool {
	matched, err := regexp.MatchString(pattern, s)
	if err == nil {
		return matched
	}
	return false
}
