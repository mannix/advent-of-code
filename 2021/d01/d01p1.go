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

	// Slice by whitespace
	sMeasurements := strings.Fields(string(input))

	// Map the strings to ints
	var measurements []int
	for _, s := range sMeasurements {
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		measurements = append(measurements, value)
	}

	// Individual increases
	individualIncreases := 0
	for i, reading := range measurements {
		// First reading is skipped
		if i > 0 {
			if reading > measurements[i-1] {
				individualIncreases++
			}
		}
	}

	// Group increases
	groupIncreases := 0
	for i := 0; i < len(measurements); i++ {
		// Skip until we have two groups to compare (4 numbers)
		if i > 2 {
			currentGroup := measurements[i] + measurements[i-1] + measurements[i-2]
			previousGroup := measurements[i-1] + measurements[i-2] + measurements[i-3]
			if currentGroup > previousGroup {
				groupIncreases++
			}
		}
	}

	fmt.Printf("Number of individual increases: %v\n", individualIncreases) // 1400
	fmt.Printf("Number of group increases: %v", groupIncreases)             // 1429
}
