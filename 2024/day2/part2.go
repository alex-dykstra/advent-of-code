package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(puzzleInputPath string) {
	file, err := os.Open(puzzleInputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLine2(line)
	}

	fmt.Printf("%v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine2(line string) int {
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
	if reprocessReportSafe(report) {
		return 1
	}
	return 0
}

func reprocessReportSafe(report []int) bool {
	for index := range report {
		reportCopy := append([]int(nil), report...)
		trimmed := append(reportCopy[:index], reportCopy[index+1:]...)
		if isSafe(trimmed) {
			return true
		}
	}
	return false
}
