package pkg

import "strconv"

func MustAtoI(data string) int {
	parsed, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}
	return parsed
}
