package issubsequence

func isSubsequence(s string, t string) bool {
	var si, ti int

	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
			ti++
		} else {
			ti++
		}
	}
	return si == len(s)
}
