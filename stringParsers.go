package stringParsers

import "strings"

func SplitOptions(s string) []string {
	return strings.Split(s, " ")
}

func Bracket(s string) string {
	return "-->" + s + "<--"
}

func BracketAll(s []string) []string {
	as := []string{}
	for _, t := range s {
		as = append(as, Bracket(t))
	}
	return as
}
