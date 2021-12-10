package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reduceOxyGenList(input []string, pos int) []string {
	match := "0"
	if numberOfOnes(input, pos) >= float32(len(input))/2.0 {
		match = "1"
	}
	return getTheKeepers(input, pos, match)
}

func reduceScrubberRating(input []string, pos int) []string {
	match := "0"
	if numberOfOnes(input, pos) < float32(len(input))/2.0 {
		match = "1"
	}
	return getTheKeepers(input, pos, match)
}

func getTheKeepers(input []string, pos int, match string) []string {
	keepers := []string{}
	for _, x := range input {
		if x != "" {
			xRune := []rune(x)
			if string(xRune[pos]) == match {
				keepers = append(keepers, x)
			}
		}
	}
	return keepers
}

func numberOfOnes(input []string, position int) float32 {
	oneCount := 0
	for _, binaryNum := range input {
		if binaryNum != "" {
			bnRune := []rune(binaryNum)
			if string(bnRune[position]) == "1" {
				oneCount += 1
			}
		}
	}
	return float32(oneCount)
}

func main() {
	if input, err := os.ReadFile("input.txt"); err == nil {
		binaryNumbers := strings.Split(string(input), "\n")
		ogr := binaryNumbers
		for i := 0; i < 12; i++ {
			ogr = reduceOxyGenList(ogr, i)
			if len(ogr) == 1 {
				break
			}
		}

		csr := binaryNumbers
		for i := 0; i < 12; i++ {
			csr = reduceScrubberRating(csr, i)
			if len(csr) == 1 {
				break
			}
		}

		if oxRating, err1 := strconv.ParseInt(ogr[0], 2, 64); err1 == nil {
			if co2Rating, err2 := strconv.ParseInt(csr[0], 2, 64); err2 == nil {
				fmt.Println(oxRating * co2Rating)
			}
		}
	}
}
