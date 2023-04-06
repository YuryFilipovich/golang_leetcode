func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, item := range operations {
		if item == "X++" || item == "++X" {
			x++
		}
		if item == "X--" || item == "--X" {
			x--
		}
	}
	return x
}