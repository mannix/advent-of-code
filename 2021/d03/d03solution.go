package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if input, err := os.ReadFile("input.txt"); err == nil {
		binaryNumbers := strings.Split(string(input), "\n")
		noOfNumbers := len(binaryNumbers) - 1 // Last line is blank
		counts := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		var gamma strings.Builder
		var epsilon strings.Builder

		for _, biNum := range binaryNumbers {
			for i, n := range biNum {
				if string(n) == "1" {
					counts[i] = counts[i] + 1
				}
			}
		}

		for _, posCnt := range counts {
			if posCnt > noOfNumbers/2 {
				gamma.WriteString("1")
				epsilon.WriteString("0")
			} else {
				gamma.WriteString("0")
				epsilon.WriteString("1")
			}
		}

		if gammaNumber, err1 := strconv.ParseInt(gamma.String(), 2, 64); err1 == nil {
			if epsilonNumber, err2 := strconv.ParseInt(epsilon.String(), 2, 64); err2 == nil {
				fmt.Println(gammaNumber * epsilonNumber)
			}
		}
	}
}
