package stringParsers

func SplitOptions(s string) []string {
	t := s
	tok := ""
	as := []string{}
	for tok, t = quoted(t); string(t) != ""; tok, t = quoted(t) {
		as = append(as, tok)
	}
	if tok != "" {
		as = append(as, tok)
	}
	return as
}

func Bracket(s string) string {
	return "-->" + s + "<--"
}

func BracketAll(s []string) []string {
	as := []string{}
	for _, t := range s {
		as = append(as, Bracket(t))
		//time.Sleep(5*time.Second)
	}
	return as
}

func quoted(s string) (string, string) {
	at := 0
	token := ""
	dlim := ""
	// skip leading blanks
	for ; string(s[at]) == " "; at++ {
		//log.Println("Stripping leading blanks:",Bracket(string(s)))
	}
	// if first char is quote, consume until ending quote, else consume until blank
	//log.Println("First char:",string(s[at]))
	if string(s[at]) == "\"" || string(s[at]) == "'" {
		dlim = string(s[at])
		at++
		esc := false
		//log.Println("a",dlim)
		for i, t := range s[at:] {
			//log.Println("b")
			if esc {
				//log.Println("c")
				esc = false
				token += string(t)
			} else if string(t) == dlim {
				//log.Println("d",token,at+i)
				return token, s[at+i+1:]
			} else {
				//log.Println("e")
				if string(t) == "\\" {
					//log.Println("f")
					esc = true
				} else {
					//log.Println("g",token)
					token += string(t)
				}
			}
			//log.Println("In a")
			//time.Sleep(5*time.Second)
		}

	} else {
		//log.Println("h",at)
		for token = ""; at < len(s); at++ {
			//log.Println("In b")
			//time.Sleep(5*time.Second)
			if string(s[at]) == " " {
				return token, s[at:]
			}
			token += string(s[at])
		}
	}
	return token, s[at:]
}
