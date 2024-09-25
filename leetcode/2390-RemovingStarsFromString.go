package leetcode

func removeStars(s string) string {
	ss := s
	for i := 0; i < len(ss); i++ {
		if ss[i]=='*' {
			if len(ss)==1 {
				return ""
			}
			if i==len(ss)-1 {
				return ss[:len(ss)-2]
			}			
			if i==0  {
				ss = ss[1:]
			}else{
				ss = ss[:i-1]+ss[i+1:]
                i--
			}
            i--
		}
	}
	return ss
}

func removeStarsPlus(s string) string {
	stack := []rune{}

	for _, ch := range s {
		if ch == '*' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, ch)
		}
	}

	return string(stack) 
}