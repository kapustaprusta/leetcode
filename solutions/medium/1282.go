package medium

func groupThePeople(groupSizes []int) [][]int {
	groupSizesMap := make(map[int][]int, 0)
	for userIdx, groupSize := range groupSizes {
		if _, isOk := groupSizesMap[groupSize]; !isOk {
			groupSizesMap[groupSize] = make([]int, 1)
			groupSizesMap[groupSize][0] = userIdx
		} else {
			groupSizesMap[groupSize] = append(groupSizesMap[groupSize], userIdx)
		}
	}

	output := make([][]int, 0)
	for groupSize, users := range groupSizesMap {
		usersCount := 0
		for usersCount < len(users) {
			output = append(output, users[usersCount:usersCount+groupSize])
			usersCount += groupSize
		}
	}

	return output
}
