package easy

func decompressRLElist(nums []int) []int {
	resNums := []int{}
	for numIdx := 0; numIdx < len(nums); numIdx += 2 {
		for resNumIdx := 0; resNumIdx < nums[numIdx]; resNumIdx++ {
			resNums = append(resNums, nums[numIdx+1])
		}
	}

	return resNums
}
