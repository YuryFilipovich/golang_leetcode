func subtractProductAndSum(n int) int {
	sum := 0
	mult := 1
	for n > 0 {
		digit := n % 10
		n = n / 10
		sum += digit
		mult *= digit
	}
	return mult - sum
}