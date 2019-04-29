package stringParsers

import (
	"log"
	"strings"
)

func SplitOptions(s string) []string {
	t := s
	tok := ""
	as := []string{}
	for ; string(t) != ""; tok, t = quoted(t) {
		as = append(as, tok)
	}
	log.Println(as)
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

func quoted(s string) (string, string) {
	at := 0
	token := ""
	// skip leading blanks
	for ; string(s[at]) != " "; at++ {
	}
	// if first char is quote, consume until ending quote, else consume until blank
	if string(s[at]) == "\"" {
		at++
		esc := false
		for _, t := range s[at:] {
			if esc {
				esc = false
				token += string(t)
			} else if string(t) == "\"" {
				return token, s[at:]
			} else {
				if string(t) == "\\" {
					esc = true
				} else {
					token += string(t)
				}
			}
		}

	} else {
		for ; string(s[at]) != " " && at < len(s); at++ {

		}
	}
	return token, s[at:]
}
