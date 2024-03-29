package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if input, err := os.ReadFile("input.txt"); err == nil {
		instructions := strings.Split(string(input), "\n")
		hor := 0
		dep := 0
		aim := 0
		for i, instruction := range instructions {
			if i == len(instructions)-1 {
				break
			}
			parts := strings.Split(instruction, " ")
			if amt, err := strconv.Atoi(parts[1]); err == nil {
				switch parts[0] {
				case "up":
					aim -= amt
				case "down":
					aim += amt
				case "forward":
					hor += amt
					dep += aim * amt
				}
			}
		}
		fmt.Println(hor * dep)
	}
}
