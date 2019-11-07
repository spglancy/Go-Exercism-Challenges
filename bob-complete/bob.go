// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	str "strings"
	uni "unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = str.Join(str.Fields(remark), " ")
	var yell bool
	var question bool
	var checked bool = false
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	for _, char := range remark{
		if uni.IsLower(char){
			yell = false
			checked = true
			break
		} else if uni.IsUpper(char) && !checked {
			checked = true
			yell = true
		}
	}
	if str.HasSuffix(remark, "?") {
		question = true
	}
	switch {
	case question && yell:
		return "Calm down, I know what I'm doing!"
	case question:
		return "Sure."
	case yell:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
