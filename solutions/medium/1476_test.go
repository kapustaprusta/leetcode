package medium

import (
	"fmt"
	"testing"
)

func TestSubrectangleQueries(t *testing.T) {
	s := Constructor([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	// 1 2 1
	// 4 3 4
	// 3 2 1
	// 1 1 1
	fmt.Println(s.GetValue(0, 2)) // return 1
	s.UpdateSubrectangle(0, 0, 3, 2, 5)
	// After this update the rectangle looks like:
	// 5 5 5
	// 5 5 5
	// 5 5 5
	// 5 5 5
	fmt.Println(s.GetValue(0, 2)) // return 5
	fmt.Println(s.GetValue(3, 1)) // return 5
	s.UpdateSubrectangle(3, 0, 3, 2, 10)
	// After this update the rectangle looks like:
	// 5   5   5
	// 5   5   5
	// 5   5   5
	// 10  10  10
	fmt.Println(s.GetValue(3, 1)) // return 10
	fmt.Println(s.GetValue(0, 2)) // return 5
}
