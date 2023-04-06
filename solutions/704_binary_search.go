func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func search(nums []int, target int) int {

	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func search(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}