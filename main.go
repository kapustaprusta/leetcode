package main

import (
	"fmt"

	"github.com/kapustaprusta/leetcode/solutions"
)

func main() {
	nums := []int{12, 1, 12, 1}
	fmt.Println(solutions.Shuffle(nums, 2))
}
