package easy

func numIdenticalPairs(nums []int) int {
	goodPairs := 0
	for outerNumIdx := 0; outerNumIdx < len(nums); outerNumIdx++ {
		for innerNumIdx := outerNumIdx + 1; innerNumIdx < len(nums); innerNumIdx++ {
			if nums[outerNumIdx] == nums[innerNumIdx] {
				goodPairs++
			}
		}
	}

	return goodPairs
}
