package easy

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prevVal := target[index[i]]
		for j := index[i] + 1; j < len(nums); j++ {
			target[j], prevVal = prevVal, target[j]
		}

		target[index[i]] = nums[i]
	}

	return target
}
