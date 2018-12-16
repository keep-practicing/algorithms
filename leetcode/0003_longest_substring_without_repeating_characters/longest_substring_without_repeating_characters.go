package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {

	var (
		start              = 0
		res                = 0
		lastOccurredRecord = make(map[rune]int)
	)

	for i, ch := range []rune(s) {
		if lastIndex, ok := lastOccurredRecord[ch]; ok && lastIndex >= start {
			start = lastIndex + 1
		}
		if i-start+1 > res {
			res = i - start + 1
		}
		lastOccurredRecord[ch] = i
	}
	return res
}
