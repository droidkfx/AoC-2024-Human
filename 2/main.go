package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reports := getData("./2/input.txt")

	start := time.Now()
	res := calculateSafeReports(reports, 0)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := calculateSafeReports(reports, 1)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

func calculateSafeReports(reports [][]int, dampenLevel int) int {
	var acc int
	for i := 0; i < len(reports); i++ {
		if reportSafe(reports[i], dampenLevel) {
			acc++
		}
	}
	return acc
}

func reportSafe(report []int, dampenLevel int) bool {
	if len(report) < 2 {
		return false
	}

	ascending := report[0] < report[1]
	var failures int
	for i := 1; i < len(report); i++ {
		delta := report[i] - report[i-1]
		var fail bool

		if delta < 0 {
			if ascending { // the report is not totally ascending
				fail = true
			}
			delta *= -1
		} else if !ascending { // the report is not totally descending
			fail = true
		}
		if delta < 1 || delta > 3 { // not gradual enough
			fail = true
		}

		if fail {
			failures++
			if failures > dampenLevel {
				return false
			} else { // try to make a safe report
				// remove left left element
				if i > 1 {
					trialReport0 := make([]int, 0, len(report))
					trialReport0 = append(trialReport0, report[:i-2]...)
					trialReport0 = append(trialReport0, report[i-1:]...)
					if reportSafe(trialReport0, dampenLevel-1) {
						return true
					}
				}

				// remove center left element
				trialReport1 := make([]int, 0, len(report))
				if i > 1 {
					trialReport1 = append(trialReport1, report[:i-1]...)
				}
				trialReport1 = append(trialReport1, report[i:]...)
				if reportSafe(trialReport1, dampenLevel-1) {
					return true
				}

				// remove center right element
				trialReport2 := make([]int, 0, len(report))
				trialReport2 = append(trialReport2, report[:i]...)
				trialReport2 = append(trialReport2, report[i+1:]...)
				if reportSafe(trialReport2, dampenLevel-1) {
					return true
				}

				// remove right right element
				if len(report) > i+2 {
					trialReport3 := make([]int, 0, len(report))
					trialReport3 = append(trialReport3, report[:i+1]...)
					trialReport3 = append(trialReport3, report[i+2:]...)
					if reportSafe(trialReport3, dampenLevel-1) {
						return true
					}
				}

				return false // could not make the report safe
			}
		}
	}
	return true // base case
}

func getData(filename string) [][]int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input := bufio.NewScanner(f)

	data := make([][]int, 0, 1000)
	for input.Scan() {
		line := strings.Split(input.Text(), " ")
		dataLine := make([]int, 0, len(line))
		for i := 0; i < len(line); i++ {
			res, err := strconv.Atoi(line[i])
			if err != nil {
				panic(err)
			}
			dataLine = append(dataLine, res)
		}
		data = append(data, dataLine)
	}

	return data
}
