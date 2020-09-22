package easy

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				result[i]++
			} else if nums[i] < nums[j] {
				result[j]++
			}
		}
	}

	return result
}
