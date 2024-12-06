package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	data := getData("./4/input.txt")

	start := time.Now()
	res := scanInput(data, countXmasInstance)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := scanInput(data, countXmasCross)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

func scanInput(data [][]byte, counter func(int, int, [][]byte) int) int {
	found := make([][]bool, len(data))
	for i := 0; i < len(data); i++ {
		found[i] = make([]bool, len(data[i]))
	}

	var acc int
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if found[i][j] {
				continue
			}

			acc += counter(i, j, data)
		}
	}
	return acc
}

func countXmasCross(i, j int, data [][]byte) int {
	if data[i][j] != 'A' {
		return 0
	}

	// cross top left -> bot right
	if (safeAccess(i-1, j-1, data) == 'M' && safeAccess(i+1, j+1, data) == 'S') || // forward
		(safeAccess(i-1, j-1, data) == 'S' && safeAccess(i+1, j+1, data) == 'M') { // backward
		// cross top right -> bot left
		if (safeAccess(i-1, j+1, data) == 'M' && safeAccess(i+1, j-1, data) == 'S') || // forward
			(safeAccess(i-1, j+1, data) == 'S' && safeAccess(i+1, j-1, data) == 'M') { // backward
			return 1
		}
	}

	return 0
}

func countXmasInstance(i, j int, data [][]byte) int {
	if data[i][j] != 'X' {
		return 0
	}
	var acc int

	// forward
	if safeAccess(i, j+1, data) == 'M' && safeAccess(i, j+2, data) == 'A' && safeAccess(i, j+3, data) == 'S' {
		acc++
	}

	// backward
	if safeAccess(i, j-1, data) == 'M' && safeAccess(i, j-2, data) == 'A' && safeAccess(i, j-3, data) == 'S' {
		acc++
	}

	// up
	if safeAccess(i+1, j, data) == 'M' && safeAccess(i+2, j, data) == 'A' && safeAccess(i+3, j, data) == 'S' {
		acc++
	}

	// down
	if safeAccess(i-1, j, data) == 'M' && safeAccess(i-2, j, data) == 'A' && safeAccess(i-3, j, data) == 'S' {
		acc++
	}

	// diag top left -> bottom right
	if safeAccess(i+1, j+1, data) == 'M' && safeAccess(i+2, j+2, data) == 'A' && safeAccess(i+3, j+3, data) == 'S' {
		acc++
	}

	// diag bottom right -> top left
	if safeAccess(i-1, j-1, data) == 'M' && safeAccess(i-2, j-2, data) == 'A' && safeAccess(i-3, j-3, data) == 'S' {
		acc++
	}

	// diag top right -> bottom left
	if safeAccess(i+1, j-1, data) == 'M' && safeAccess(i+2, j-2, data) == 'A' && safeAccess(i+3, j-3, data) == 'S' {
		acc++
	}

	// diag bottom left -> top right
	if safeAccess(i-1, j+1, data) == 'M' && safeAccess(i-2, j+2, data) == 'A' && safeAccess(i-3, j+3, data) == 'S' {
		acc++
	}

	return acc
}

func safeAccess(i, j int, data [][]byte) byte {
	if i < 0 || i >= len(data) {
		return 0
	}
	if j < 0 || j >= len(data[i]) {
		return 0
	}
	return data[i][j]
}

func getData(filename string) [][]byte {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	input := bufio.NewScanner(f)

	data := make([][]byte, 0, 1000)
	for input.Scan() {
		line := input.Text()
		data = append(data, []byte(line))
	}

	return data
}
