// Package isogram checks if string is an isogram
package isogram

import "strings"

// IsIsogram checks if string is an isogram
func IsIsogram(in string) bool {
  if len(in) <= 1 {
    return true
  }
  str := strings.ToLower(in)
  str = strings.Replace(str, " ", "", -1)
  str = strings.Replace(str, "-", "", -1)
  return Contains(str[1:len(str)], string(str[0]))
}

// Contains recursively check if target string exists in the input
func Contains(in string, target string) bool{
  if len(in) <= 1 {
    return true
  }
  if strings.Contains(in, target) {
    return false
  }
  return Contains(in[1:len(in)], string(in[0]))
}
