package main

func isValid(s string) bool {
	left := []rune{} // 模拟一个栈
	for _, c := range s {
		// 如果是左括号
		if c == '(' || c == '{' || c == '[' {
			left = append(left, c)
		} else {
			if len(left) != 0 && leftOf(c) == left[len(left)-1] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		}
	}
	return len(left) == 0
}

func leftOf(c rune) rune {
	switch c {
	case '}':
		return '{'
	case ')':
		return '('
	default:
		return '['
	}
}

func main() {
	s := "({[]})"
	println(isValid(s))
}
