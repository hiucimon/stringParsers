package stringParsers

import "strings"

func SplitOptions(s string) []string {
	return strings.Split(s, " ")
}

func bracket(s string) string {
	return "-->" + s + "<--"
}
