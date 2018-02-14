// Package twofer generates short message
package twofer

import "strings"

var msg = "One for you, one for me."

// ShareWith generates short message.
func ShareWith(name string) string {
	if name != "" {
		return strings.Replace(msg, "you", name, -1)
	}
	return msg
}
