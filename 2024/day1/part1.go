package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(puzzleInputPath string) {
	file, err := os.Open(puzzleInputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			log.Fatal("Split didn't work")
		}
		if leftValue, err := strconv.Atoi(split[0]); err == nil {
			left = append(left, leftValue)
		}
		if rightValue, err := strconv.Atoi(split[1]); err == nil {
			right = append(right, rightValue)
		}
	}

	if len(left) != len(right) {
		log.Fatal("Left and right lists are not the same length!")
	}

	slices.Sort(left)
	slices.Sort(right)

	var distanceDifferences []int
	for index, leftElement := range left {
		rightElement := right[index]
		difference := absInt(leftElement, rightElement)
		distanceDifferences = append(distanceDifferences, difference)
	}

	total := listSum(distanceDifferences)

	fmt.Printf("%v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func listSum(list []int) int {
	total := 0
	for _, difference := range list {
		total += difference
	}
	return total
}
