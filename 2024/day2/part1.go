package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(puzzleInputPath string) {
	file, err := os.Open(puzzleInputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLine(line)
	}

	fmt.Printf("%v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine(line string) int {
	var report []int
	split := strings.Split(line, " ")
	for _, splitItem := range split {
		if level, err := strconv.Atoi(splitItem); err == nil {
			report = append(report, level)
		}
	}
	if isSafe(report) {
		return 1
	}
	return 0
}

func isSafe(values []int) bool {
	if isIncreasing(values) || isDecreasing(values) {
		return true
	}
	return false
}

func isIncreasing(values []int) bool {
	previousValue := values[0]
	for _, value := range values[1:] {
		if value > previousValue {
			if value-previousValue > 3 {
				return false
			}
		} else {
			return false
		}
		previousValue = value
	}
	return true
}

func isDecreasing(values []int) bool {
	previousValue := values[0]
	for _, value := range values[1:] {
		if value < previousValue {
			if previousValue-value > 3 {
				return false
			}
		} else {
			return false
		}
		previousValue = value
	}
	return true
}
