package raindrops

import (
	"strconv"
	"strings"
)

func Convert(i int) string {
	var output []string
	if i%3 == 0 {
		output = append(output, "Pling")
	}
	if i%5 == 0 {
		output = append(output, "Plang")
	}
	if i%7 == 0 {
		output = append(output, "Plong")
	}
	if len(output) == 0 {
		return strconv.Itoa(i)
	}
	return strings.Join(output, "")
}
