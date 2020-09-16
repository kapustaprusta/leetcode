package easy

func numJewelsInStones(J string, S string) int {
	realJewelsNumber := 0
	jewelsMap := make(map[rune]bool)

	for _, currJewel := range J {
		jewelsMap[currJewel] = true
	}

	for _, currStone := range S {
		if _, isExist := jewelsMap[currStone]; isExist {
			realJewelsNumber++
		}
	}

	return realJewelsNumber
}
