package easy

func numberOfSteps(num int) int {
	totalSteps := 0
	for num != 0 {
		if num&1 == 1 {
			num--
			totalSteps++
			if num == 0 {
				break
			}
		}

		num = num >> 1
		totalSteps++
	}

	return totalSteps
}
