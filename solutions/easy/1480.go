package easy

// RunningSum ...
func RunningSum(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]
	for numIdx, num := range nums[1:] {
		result[numIdx+1] = result[numIdx] + num
	}

	return result
}
