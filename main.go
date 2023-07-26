package main

var m = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isOpen(c byte) bool {
	return c == '(' || c == '{' || c == '['
}

func isValid(s string) bool {
	stack := make([]int32, 0)
	st := 0

	for _, p := range s {
		if isOpen(byte(p)) {
			st += 1
			stack = append(stack, p)
		} else {
			st -= 1
			if len(stack) == 0 {
				return false
			}

			lastBracket := stack[len(stack)-1]
			if m[byte(lastBracket)] != byte(p) {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if st < 0 {
			return false
		}
	}

	return (st == 0)
}
