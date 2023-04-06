func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0

	for _, num := range nums {
		if num == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
	}

	return maxCount
}