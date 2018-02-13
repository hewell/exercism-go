// Package acronym generates acronyms
package acronym

import "regexp"
import "strings"

// Abbreviate generates acronym from input.
func Abbreviate(s string) string {
	re := regexp.MustCompile("\\b[a-zA-Z]")
	res := re.FindAllString(s, -1)
	acronym := strings.Join(res, "")
	return strings.ToUpper(acronym)
}
