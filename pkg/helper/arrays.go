package helper

func SafeAccess[T any](i, j int, data [][]T) T {
	res, _ := SafeAccessOption(i, j, data)
	return res
}

func SafeAccessOption[T any](i, j int, data [][]T) (T, bool) {
	var zeroValue T
	if i < 0 || i >= len(data) {
		return zeroValue, false
	}
	if j < 0 || j >= len(data[i]) {
		return zeroValue, false
	}
	return data[i][j], true
}
