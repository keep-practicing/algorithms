package validparentheses

func isValid(s string) bool {
	var (
		stacks  []rune
		mapping = map[rune]rune{')': '(', ']': '[', '}': '{'}
	)
	for _, char := range s {
		if char == ')' || char == '}' || char == ']' {
			var topElement rune
			if len(stacks) > 0 {
				topElement = stacks[len(stacks)-1]
				stacks = stacks[:len(stacks)-1]
			} else {
				topElement = '#'
			}
			if mapping[char] != topElement {
				return false
			}
		} else {
			stacks = append(stacks, char)
		}
	}
	if len(stacks) > 0 {
		return false
	}
	return true
}
