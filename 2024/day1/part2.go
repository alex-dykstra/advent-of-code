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

	var left []int
	right := make(map[int]int)

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
			right[rightValue] = right[rightValue] + 1
		}
	}

	var similarityScores []int
	for _, leftElement := range left {
		similarityScore := leftElement * right[leftElement]
		similarityScores = append(similarityScores, similarityScore)
	}

	total := listSum(similarityScores)
	fmt.Printf("%v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
