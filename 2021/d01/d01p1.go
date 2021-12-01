package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Get data from input file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Cast to string
	strInput := string(input)

	// Slice by whitespace
	sMeasurements := strings.Fields(strInput)

	// Map the strings to ints
	var measurements []int
	for _, s := range sMeasurements {
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		measurements = append(measurements, value)
	}

	// Count increases
	increases := 0
	for i, reading := range measurements {
		// First reading is skipped
		if i > 0 {
			if reading > measurements[i-1] {
				increases++
			}
		}
	}

	fmt.Printf("Number of increases: %v", increases) // 1400
}
