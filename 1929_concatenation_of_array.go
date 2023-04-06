func getConcatenation(nums []int) []int {
	newSlice := append(nums, nums...)
	return newSlice
}