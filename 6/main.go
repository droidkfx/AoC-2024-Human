package main

import (
	"adventofcode2024/pkg/helper"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./6/input.txt")
	if err != nil {
		panic(err)
	}
	guardMap, guard := getData(f)
	_ = f.Close()

	start := time.Now()
	res, _ := countGuardSpaces(guardMap, guard)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := countPossibleLoops(guardMap, guard)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

// this is brute force nonsense. I have might be able to do better by restricting to locations where we have visited
// at least once. That would cut down the search space significantly. I think there might even be a far more
// effecient solution out there that would require using a graph of paths instead and finding a loop? unsure.
// a 2s run time is pathetic though.
func countPossibleLoops(guardMap [][]bool, guard guardPosition) int {
	var acc int
	for j := 0; j < len(guardMap); j++ {
		for i := 0; i < len(guardMap[j]); i++ {
			if !guardMap[j][i] && !(j == guard.y && i == guard.x) {
				guardMap[j][i] = true
				_, loop := countGuardSpaces(guardMap, guard)
				if loop {
					acc++
				}
				guardMap[j][i] = false
			}
		}
	}
	return acc
}

func countGuardSpaces(guardMap [][]bool, guard guardPosition) (int, bool) {
	// assume the guard is on the map someplace
	acc := 1
	visited := make([][][]bool, 0, len(guardMap))
	for _, _ = range guardMap {
		visitLine := make([][]bool, len(guardMap[0]))
		for i, _ := range visitLine {
			// 4 directions plus 1 ignored
			visitLine[i] = make([]bool, 5)
		}
		visited = append(visited, visitLine)
	}
	visited[guard.y][guard.x][guard.d] = true

	for guard.x >= 0 && guard.x < len(guardMap[0]) && guard.y >= 0 && guard.y < len(guardMap) {
		var nextX int
		var nextY int
		switch guard.d {
		case Dir_UP:
			nextX = guard.x
			nextY = guard.y - 1
		case Dir_DOWN:
			nextX = guard.x
			nextY = guard.y + 1
		case Dir_LEFT:
			nextX = guard.x - 1
			nextY = guard.y
		case Dir_RIGHT:
			nextX = guard.x + 1
			nextY = guard.y
		}

		v, ok := helper.SafeAccessOption(nextY, nextX, guardMap)
		if ok && v {
			switch guard.d {
			case Dir_UP:
				guard.d = Dir_RIGHT
			case Dir_DOWN:
				guard.d = Dir_LEFT
			case Dir_LEFT:
				guard.d = Dir_UP
			case Dir_RIGHT:
				guard.d = Dir_DOWN
			}
		} else {
			if ok {
				if visited[nextY][nextX][guard.d] { // loop detected, we have already been to the next space in the same direction
					return acc, true
				}

				// no loop but have we ever been here?
				var seen bool
				for i := 0; !seen && i < len(visited[nextY][nextX]); i++ {
					seen = visited[nextY][nextX][i]
				}
				if !seen {
					visited[nextY][nextX][guard.d] = true
					acc++
				}
			}
			guard.x = nextX
			guard.y = nextY
		}
	}

	return acc, false
}

type direction int

const (
	_ direction = iota
	Dir_UP
	Dir_DOWN
	Dir_LEFT
	Dir_RIGHT
)

type guardPosition struct {
	x int
	y int
	d direction
}

func getData(f io.Reader) ([][]bool, guardPosition) {
	input := bufio.NewScanner(f)

	guardMap := make([][]bool, 0, 1000)
	var guardX int
	var guardY int
	var guardDirection direction
	for input.Scan() {
		line := input.Text()
		lineMap := make([]bool, len(line))
		for i, c := range line {
			if c == '#' {
				lineMap[i] = true
			}
			if c != '.' && c != '#' {
				guardX = i
				guardY = len(guardMap)
				switch c {
				case '^':
					guardDirection = Dir_UP
				case '>':
					guardDirection = Dir_RIGHT
				case '<':
					guardDirection = Dir_LEFT
				case 'V':
					guardDirection = Dir_DOWN
				}
			}

		}
		guardMap = append(guardMap, lineMap)
	}

	return guardMap, guardPosition{x: guardX, y: guardY, d: guardDirection}
}
