package medium

import "sync"

var (
	wg sync.WaitGroup
)

// SubrectangleQueries ...
type SubrectangleQueries struct {
	vals [][]int
}

// Constructor ...
func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

// UpdateSubrectangle ...
func (s *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for rowIdx := row1; rowIdx <= row2; rowIdx++ {
		wg.Add(1)
		go func(rowIdx int) {
			for colIdx := col1; colIdx <= col2; colIdx++ {
				s.vals[rowIdx][colIdx] = newValue
			}
			wg.Done()
		}(rowIdx)
	}

	wg.Wait()
}

// GetValue ...
func (s *SubrectangleQueries) GetValue(row int, col int) int {
	return s.vals[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
