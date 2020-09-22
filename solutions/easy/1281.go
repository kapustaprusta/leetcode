package easy

func subtractProductAndSum(n int) int {
	s := 0
	p := 1

	for {
		d := n - (n/10)*10
		s += d
		p *= d
		n /= 10
		if n == 0 {
			break
		}
	}

	return p - s
}
