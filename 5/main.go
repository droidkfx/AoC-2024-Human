package main

import (
	"adventofcode2024/pkg/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	rules, orders := getData("./5/input.txt")

	start := time.Now()
	res := countMiddleOfCorrectOrders(rules, orders)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := countMiddleOfCorrectedOrders(rules, orders)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

func countMiddleOfCorrectedOrders(rules map[int][]int, orders [][]int) int {
	var acc int

	for _, order := range orders {
		firstSeen := getFirstSeenMap(order)

		// I am not proud of this.... This is brute force nonsense. I think I could just quicksort it, but I would have
		// to sort out how to determine order. Not hard I think I could just use the map as the function, but I am done
		// for tonight and I leave that as an exercise to the reader.
		/// something like:
		// func less(a,b int) int {
		//   return rules[a] - rules[b] <- not exactly because rules is a map of int -> []int. maybe for loop it?
		// }
		if !doesOrderMeetRules(rules, firstSeen) {
			for i := 0; !doesOrderMeetRules(rules, firstSeen); i++ {
				correctOrdering(rules, order, firstSeen)
				firstSeen = getFirstSeenMap(order)
				if i > len(orders) { // this shouldn't happen right? Because this is essentially a swap sort?
					panic("non quadratic solution")
				}
			}
			acc += order[(len(order)-1)/2]
		}
	}

	return acc
}

func correctOrdering(rules map[int][]int, order []int, firstSeen map[int]int) {
	for k, v := range rules {
		firstNumbSeenIdx, ok := firstSeen[k]
		if !ok {
			continue
		}

		for _, n := range v {
			secondNumbSeenIdx, ok := firstSeen[n]
			if !ok {
				continue
			}

			if firstNumbSeenIdx > secondNumbSeenIdx {
				order[firstNumbSeenIdx], order[secondNumbSeenIdx] = order[secondNumbSeenIdx], order[firstNumbSeenIdx]
				firstSeen[k] = secondNumbSeenIdx
				firstSeen[n] = firstNumbSeenIdx
			}
		}
	}
}

func countMiddleOfCorrectOrders(rules map[int][]int, orders [][]int) int {
	var acc int

	for _, order := range orders {
		firstSeen := getFirstSeenMap(order)

		if doesOrderMeetRules(rules, firstSeen) {
			acc += order[(len(order)-1)/2]
		}
	}

	return acc
}

func getFirstSeenMap(order []int) map[int]int {
	firstSeen := make(map[int]int)
	for i, page := range order {
		_, ok := firstSeen[page]
		if !ok {
			firstSeen[page] = i
		}
	}
	return firstSeen
}

func doesOrderMeetRules(rules map[int][]int, firstSeen map[int]int) bool {
	for k, v := range rules {
		firstNumbSeenIdx, ok := firstSeen[k]
		if !ok {
			continue
		}

		for _, n := range v {
			secondNumbSeenIdx, ok := firstSeen[n]
			if !ok {
				continue
			}

			if firstNumbSeenIdx > secondNumbSeenIdx {
				return false
			}
		}
	}

	return true
}

func getData(filename string) (map[int][]int, [][]int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	input := bufio.NewScanner(f)

	rules := make(map[int][]int)
	reports := make([][]int, 0, 1000)
	var rulesDone bool
	for input.Scan() {
		line := input.Text()
		if !rulesDone && line == "" {
			rulesDone = true
			continue
		}

		if !rulesDone {
			ruleNumbers := strings.Split(line, "|")
			firstNumb := helper.MustAtoI(ruleNumbers[0])
			secondNumb := helper.MustAtoI(ruleNumbers[1])

			v, ok := rules[firstNumb]
			if !ok {
				v = make([]int, 0)
			}
			v = append(v, secondNumb)
			rules[firstNumb] = v // assign again since append could have changed the pointer
		} else {
			pages := strings.Split(line, ",")
			newReport := make([]int, 0, len(pages))
			for _, p := range pages {
				newReport = append(newReport, helper.MustAtoI(p))
			}
			reports = append(reports, newReport)
		}
	}

	return rules, reports
}
