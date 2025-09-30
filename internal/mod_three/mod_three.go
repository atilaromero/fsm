package modthree

func ModThree(s string) int {
	return bStringToInt(s) % 3
}

// bStringToInt converts a string of 0s and 1s to an integer.
// Any non '0' or '1' characters are skipped.
func bStringToInt(s string) int {
	var n int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '0' {
			n = n << 1
		} else if c == '1' {
			n = (n << 1) + 1
		} else {
			// ignore any non-binary character
			continue
		}
	}
	return n
}
