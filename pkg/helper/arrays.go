package helper

func SafeAccess(i, j int, data [][]byte) byte {
	if i < 0 || i >= len(data) {
		return 0
	}
	if j < 0 || j >= len(data[i]) {
		return 0
	}
	return data[i][j]
}
