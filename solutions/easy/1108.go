package easy

// defangIPaddr ..
func defangIPaddr(address string) string {
	for charIdx := 0; charIdx < len(address); charIdx++ {
		if address[charIdx] == '.' {
			address = address[:charIdx] + "[.]" + address[charIdx+1:]
			charIdx += 2
		}
	}

	return address
}
