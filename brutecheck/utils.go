package main

// keywordBounds returns minimum and maximum keyword length.
func keywordBounds(keywords []string) (min, max int) {
	min, max = len(keywords[0]), len(keywords[0])
	for _, keyword := range keywords[1:] {
		n := len(keyword)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

// findUniqueChar finds first unique character from names that is safe to check.
//
// returns -1 when one does not exist.
func findUniqueChar(names []string) int {
	minlen, _ := keywordBounds(names)
next:
	for i := 0; i < minlen; i++ {
		var seen [256]bool
		for _, name := range names {
			if seen[name[i]] {
				continue next
			}
			seen[name[i]] = true
		}
		return i
	}
	return -1
}
