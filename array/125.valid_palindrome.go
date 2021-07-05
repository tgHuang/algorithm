
func isPalindrome(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}
	if sLen == 1 {
		return true
	}
	lIndex, rIndex := 0, sLen-1
	for {
		if lIndex >= rIndex {
			fmt.Println(lIndex)
			fmt.Println(rIndex)
			return true
		}
		lstr := s[lIndex]
		if lstr < '0' || ('9' < lstr && lstr < 'A') || (lstr > 'Z' && lstr < 'a') || lstr > 'z' {
			lIndex++
			continue
		}

		rstr := s[rIndex]
		if rstr < '0' || ('9' < rstr && rstr < 'A') || (rstr > 'Z' && rstr < 'a') || rstr > 'z' {
			rIndex--
			continue
		}
		if lstr == rstr {
			lIndex++
			rIndex--
			continue
		}
		if lstr > rstr && (rstr >= 'A' && rstr <= 'Z') {
			num2 := int32(lstr) - int32(rstr)
			if num2 == 'a'-'A' {
				lIndex++
				rIndex--
				continue
			}
		}
		if lstr < rstr && (lstr >= 'A' && lstr <= 'Z') {
			num2 := int32(rstr) - int32(lstr)
			if num2 == 'a'-'A' {
				lIndex++
				rIndex--
				continue
			}
		}

		return false
	}
}
