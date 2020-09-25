package easy

func xorOperation(n int, start int) int {
	xor := 0
	for numIdx := 0; numIdx < n; numIdx++ {
		xor ^= start + 2*numIdx
	}

	return xor
}
