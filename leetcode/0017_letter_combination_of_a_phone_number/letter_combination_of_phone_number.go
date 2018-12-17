package lettercombinationsofaphonenumber

var (
	letterMap = []string{
		" ",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	res []string
)

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}

	findCombinations(digits, 0, "")
	return res
}

func findCombinations(digits string, index int, s string) {
	if index == len(digits) {
		res = append(res, s)
		return
	}

	ch := digits[index]
	letters := letterMap[ch-'0']
	for _, i := range letters {
		findCombinations(digits, index+1, s+string(i))
	}
	return
}
