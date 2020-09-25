package easy

import (
	"fmt"
	"testing"
)

func TestDecompressRLElist(t *testing.T) {
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
