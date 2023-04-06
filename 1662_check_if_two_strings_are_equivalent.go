func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var w1 string
	var w2 string
	for i := 0; i < len(word1); i++ {
		// Loop over each string in the array
		for j := 0; j < len(word1[i]); j++ {
			// Print each letter
			w1 += string(word1[i][j])
		}
	}
	for i := 0; i < len(word2); i++ {
		// Loop over each string in the array
		for j := 0; j < len(word2[i]); j++ {
			// Print each letter
			w2 += string(word2[i][j])
		}
	}

	if w1 == w2 {
		return true
	}
	return false
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := strings.Join(word1, "")
	w2 := strings.Join(word2, "")
	if w1 == w2 {
		return true
	}
	return false
}