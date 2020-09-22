package medium

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax := 0
	colMax := 0

	rSkyLine := make(map[int]int, len(grid))
	tSkyLine := make(map[int]int, len(grid[0]))

	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
			if grid[rowIdx][colIdx] > rowMax {
				rowMax = grid[rowIdx][colIdx]
			}
			if grid[colIdx][rowIdx] > colMax {
				colMax = grid[colIdx][rowIdx]
			}
		}
		tSkyLine[rowIdx] = colMax
		rSkyLine[rowIdx] = rowMax

		colMax = 0
		rowMax = 0
	}

	totalSum := 0
	for rowIdx := 0; rowIdx < len(grid); rowIdx++ {
		for colIdx := 0; colIdx < len(grid[0]); colIdx++ {
			if tSkyLine[colIdx] > rSkyLine[rowIdx] {
				totalSum += rSkyLine[rowIdx] - grid[rowIdx][colIdx]
			} else {
				totalSum += tSkyLine[colIdx] - grid[rowIdx][colIdx]
			}
		}
	}

	return totalSum
}
