package main

import (
	"adventofcode2024/pkg"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	data := getData("./3/input.txt")

	start := time.Now()
	res := multiply(data, false)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(elapsed)

	start2 := time.Now()
	res2 := multiply(data, true)
	elapsed2 := time.Since(start2)
	fmt.Println(res2)
	fmt.Println(elapsed2)
}

func multiply(data string, enableMul bool) int {
	exp, err := regexp.Compile("(mul|do|don't)\\((?:(\\d{1,3}),(\\d{1,3}))?\\)")
	if err != nil {
		panic(err)
	}

	results := exp.FindAllStringSubmatch(data, -1)

	var acc int
	mulEnabled := true
	for _, res := range results {
		instruction := res[1]
		switch instruction {
		case "do":
			if enableMul {
				mulEnabled = true
			}
		case "don't":
			if enableMul {
				mulEnabled = false
			}
		case "mul":
			if mulEnabled {
				acc += pkg.MustAtoI(res[2]) * pkg.MustAtoI(res[3])
			}
		}

	}

	return acc
}

func getData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(data)
}
