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
	res := countGuardSpaces(guardMap, guard)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	//start2 := time.Now()
	//res2 := countMiddleOfCorrectedOrders(rules, orders)
	//elapsed2 := time.Since(start2)
	//fmt.Println(res2)
	//fmt.Println(elapsed2)
}

func countGuardSpaces(guardMap [][]bool, guard guardPosition) int {
	// assume the guard is on the map someplace
	acc := 1
	visited := make([][]bool, 0, len(guardMap))
	for _, _ = range guardMap {
		visited = append(visited, make([]bool, len(guardMap[0])))
	}
	visited[guard.y][guard.x] = true

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
			if ok && !visited[nextY][nextX] {
				visited[nextY][nextX] = true
				acc++
			}
			guard.x = nextX
			guard.y = nextY
		}
	}

	return acc
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
