package main

import (
	"adventofcode2024/pkg/helper"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type calibrationData struct {
	result  int
	numbers []int
}

func main() {
	f, err := os.Open("./7/input.txt")
	if err != nil {
		panic(err)
	}
	cData := getData(f)
	_ = f.Close()

	start := time.Now()
	res := sumValidCalibrationData(cData, false)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := sumValidCalibrationData(cData, true)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

func sumValidCalibrationData(dataList []calibrationData, concatActive bool) int {
	var acc int
	for _, data := range dataList {
		if isCalibrationDataPossible(data, concatActive) {
			acc += data.result
		}
	}
	return acc
}

type operator int

const (
	ADD operator = iota
	MULT
	CONCAT
)

func isCalibrationDataPossible(data calibrationData, concatActive bool) bool {
	operators := make([]operator, len(data.numbers)-1)
	for {
		acc := data.numbers[0]
		for i := 1; i < len(data.numbers); i++ {
			o := operators[i-1]
			if o == ADD {
				acc += data.numbers[i]
			} else if o == CONCAT {
				acc = helper.MustAtoI(strconv.Itoa(acc) + strconv.Itoa(data.numbers[i]))
			} else {
				acc *= data.numbers[i]
			}
		}
		if acc == data.result {
			return true
		}
		if !nextPermutation(operators, concatActive) {
			return false
		}
	}
}

func nextPermutation(operators []operator, concatActive bool) bool {
	i := 0
	for ; i < len(operators); i++ {
		o := operators[i]
		if o == ADD {
			operators[i] = MULT
			return true
		} else if concatActive && o == MULT {
			operators[i] = CONCAT
			return true
		} else {
			operators[i] = ADD
		}
	}
	return !(i == len(operators) && operators[len(operators)-1] == ADD)
}

func getData(f io.Reader) []calibrationData {
	input := bufio.NewScanner(f)

	cData := make([]calibrationData, 0, 1000)
	for input.Scan() {
		line := input.Text()
		parts := strings.Split(line, ":")
		result := helper.MustAtoI(parts[0])
		inputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		inputList := make([]int, 0, len(inputs))
		for _, i := range inputs {
			inputList = append(inputList, helper.MustAtoI(i))
		}
		cData = append(cData, calibrationData{result: result, numbers: inputList})
	}

	return cData
}
