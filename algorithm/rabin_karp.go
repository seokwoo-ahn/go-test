package algorithm

func RabinKarp(text, pattern string) int {
	textHash, patternHash := 0, 0
	powPrime, modPrime := 3, 373
	pow := 1

	for i := len(pattern) - 1; i >= 0; i-- {
		textHash = textHash + int(text[i])*pow
		patternHash = patternHash + int(pattern[i])*pow
		if i > 0 {
			pow = pow * powPrime
		}
	}

	for i := 0; i < len(text)-len(pattern); i++ {
		if i > 0 {
			textHash = (textHash-int(text[i-1])*pow)*powPrime + int(text[i+len(pattern)-1])
		}

		if textHash%modPrime == patternHash%modPrime {
			check := true
			for j := i; j < i+len(pattern); j++ {
				if text[j] != pattern[j-i] {
					check = false
					break
				}
			}
			if check {
				return i
			}
		}
	}
	return -1
}
