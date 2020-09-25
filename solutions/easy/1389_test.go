package easy

import (
	"fmt"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	fmt.Println(createTargetArray([]int{7, 6, 5, 5, 5, 4, 5, 5}, []int{0, 1, 1, 2, 4, 2, 3, 6}))
}
