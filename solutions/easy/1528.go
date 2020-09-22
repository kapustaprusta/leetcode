package easy

func restoreString(s string, indices []int) string {
	restoredStrRaw := make([]byte, len(s))
	for indiceIdx, indice := range indices {
		restoredStrRaw[indice] = s[indiceIdx]
	}

	return string(restoredStrRaw)
}
