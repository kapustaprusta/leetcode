package easy

// KidsWithCandies ...
func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandiesNum := 0
	for _, candiesNum := range candies {
		if candiesNum > maxCandiesNum {
			maxCandiesNum = candiesNum
		}
	}

	result := make([]bool, len(candies))
	for candiesIdx, candiesNum := range candies {
		result[candiesIdx] = candiesNum+extraCandies >= maxCandiesNum
	}

	return result
}
