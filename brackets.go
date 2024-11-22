package brackets

func IsBalanced(s string) bool {
	open := 0
	wild := 0
	l := len(s)
	for i := 0; i < l; i++ {
		char := s[i]
		if char == '(' {
			if i == l-1 {
				return false
			}
			open++
		} else if char == ')' {
			open--
			if open < 0 {
				if wild > 0 {
					wild--
					open++
				} else {
					return false
				}
			}
		} else if char == '*' {
			wild++
		}
	}
	return open-wild == 0 || l == wild
}
