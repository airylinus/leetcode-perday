package slide_window

func LongestSubString(inString string) (r string) {
	var h, e int
	e = 1
	for e < len(inString) {
		if inString[e] == inString[h] {
			h += 1
			e = h + 1
		} else {
			e += 1
		}
	}
	r = inString[h:e]
	return r
}
