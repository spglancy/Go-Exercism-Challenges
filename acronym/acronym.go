// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(words string) string {
	var output []rune
	r := regexp.MustCompile(`[^a-zA-Z'\s]`)
	words = r.ReplaceAllString(words, " ")
	for _, item := range strings.Fields(words) {
		output = append(output, unicode.ToUpper(rune(item[0])))
	}
	return string(output)
}
