package isogram

import "unicode"

func IsIsogram(s string) bool {
	used := make(map[rune]int)
	for _, val := range s {
		i := unicode.ToLower(val)
		if used[i] == 1 && i != '-' && i != ' ' {
			return false
		} else {
			used[i] = 1
		}
	}
	return true
}
