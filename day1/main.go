package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadWholeFile(fname string) string {
	fileBytes, err := os.ReadFile(fname)
	if err != nil {
		panic("error: could not read file " + fname)
	}
	return string(fileBytes)
}

func PartOne(calStr string) int {
	maxCal, elfCal := -1, -1
	cals := strings.Split(calStr, "\n")

	for _, cal := range cals {
		if cal == "" {
			if maxCal < elfCal {
				maxCal = elfCal
			}
			elfCal = 0
		} else {
			currCal, _ := strconv.Atoi(cal)
			elfCal += currCal
		}
	}
	return maxCal
}

func PartTwo(calStr string) int {
	maxCals := []int{-1, -1, -1}
	elfCal := -1
	cals := strings.Split(calStr, "\n")

	for _, cal := range cals {
		if cal == "" {
			if maxCals[2] < elfCal {
				maxCals[0] = maxCals[1]
				maxCals[1] = maxCals[2]
				maxCals[2] = elfCal

			} else if maxCals[1] < elfCal {
				maxCals[0] = maxCals[1]
				maxCals[1] = elfCal

			} else if maxCals[0] < elfCal {
				maxCals[0] = elfCal

			}
			elfCal = 0
		} else {
			currCal, _ := strconv.Atoi(cal)
			elfCal += currCal
		}
	}
	return maxCals[0] + maxCals[1] + maxCals[2]
}

func main() {
	fname := "input.txt"
	fcont := ReadWholeFile(fname)
	fmt.Printf("Answer: %v\n", PartTwo(fcont))
}
