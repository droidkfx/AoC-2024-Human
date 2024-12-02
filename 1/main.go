package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	l1, l2 := getData("./1/input.txt")

	start := time.Now()
	res := calculateDistance(l1, l2)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start = time.Now()
	res2 := similarityScore(l1, l2)
	elapsed = time.Since(start)
	fmt.Println(res2)
	fmt.Println(elapsed)
}

func similarityScore(l1, l2 []int) int {
	acc := 0
	righti := 0
	lastCount := 0
	for i := 0; i < len(l1); i++ {
		if i > 0 && l1[i-1] == l1[i] {
			acc += lastCount * l1[i]
		} else {
			for ; righti < len(l2) && l1[i] >= l2[righti]; righti++ {
				if l1[i] == l2[righti] {
					acc += l1[i]
					lastCount++
				}
			}
		}
	}
	return acc
}

func calculateDistance(l1, l2 []int) int {
	slices.Sort(l1)
	slices.Sort(l2)

	acc := 0
	for i := 0; i < len(l1); i++ {
		dist := l1[i] - l2[i]
		if dist < 0 {
			dist *= -1
		}
		acc += dist
	}
	return acc
}

func getData(filename string) ([]int, []int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input := bufio.NewScanner(f)

	list1, list2 := make([]int, 0, 1000), make([]int, 0, 1000)
	for input.Scan() {
		data := strings.Fields(input.Text())
		if len(data) != 2 {
			fmt.Println("can't parse:", data)
			continue
		}

		left, err := strconv.Atoi(data[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		list1 = append(list1, left)
		list2 = append(list2, right)
	}

	return list1, list2
}
