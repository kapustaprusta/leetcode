package easy

// Shuffle ...
func Shuffle(nums []int, n int) []int {
	result := make([]int, len(nums))
	for numIdx := 0; numIdx < n; numIdx++ {
		result[numIdx*2] = nums[numIdx]
		result[numIdx*2+1] = nums[n+numIdx]
	}

	return result
}
