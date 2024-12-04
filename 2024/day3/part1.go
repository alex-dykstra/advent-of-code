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

	lines := ""
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		lines += line
	}

	lines = strings.ReplaceAll(lines, "\n", "-")
	clean := cleanse(lines)
	value := processCleansed(clean)
	fmt.Printf("%v\n", value)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processCleansed(clean string) int {
	value := 0
	split := strings.Fields(clean)
	for _, compute := range split {
		// fmt.Println(compute)
		computeNaked := strings.Replace(compute, "mul(", "", -1)
		computeNaked = strings.Replace(computeNaked, ")", "", -1)
		computeSplit := strings.Split(computeNaked, ",")
		if first, err := strconv.Atoi(computeSplit[0]); err == nil {
			if second, err := strconv.Atoi(computeSplit[1]); err == nil {
				value += (first * second)
			}
		}
	}
	return value
}

func isNumber(possibleNumber string) bool {
	if strings.Contains("0123456789", possibleNumber) {
		return true
	}
	return false
}

func cleanse(input string) string {
	cleansed := ""
	currentLexicon := ""
	currentNumber := ""
	for _, tokenRune := range input {
		token := string(tokenRune)
		if currentLexicon == "" {
			if token == "m" {
				currentLexicon += token
			}
		} else {
			if currentLexicon == "m" {
				if token == "u" {
					currentLexicon += token
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
			} else if currentLexicon == "mu" {
				if token == "l" {
					currentLexicon += token
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
			} else if currentLexicon == "mul" {
				if token == "(" {
					currentLexicon += token
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
			} else if currentLexicon == "mul(" {
				if isNumber(token) && len(currentNumber) < 3 {
					currentLexicon += token
					currentNumber += token
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
			} else if isNumber(token) && strings.Contains(currentLexicon, "mul(") {
				if len(currentNumber) < 3 {
					currentLexicon += token
					currentNumber += token
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
				// fmt.Println(currentLexicon)
			} else if token == "," {
				if strings.Contains(currentLexicon, "mul(") && currentNumber != "" {
					currentLexicon += token
				} else {
					currentLexicon = ""
				}
				currentNumber = ""
			} else if token == ")" {
				// fmt.Printf("%v\n", currentLexicon)
				// fmt.Printf("%v\n", currentNumber)
				if strings.Contains(currentLexicon, "mul(") && strings.Contains(currentLexicon, ",") {
					currentLexicon += token
					cleansed += currentLexicon
					cleansed += " "
					currentLexicon = ""
					currentNumber = ""
				} else {
					currentLexicon = ""
					currentNumber = ""
				}
			} else {
				currentLexicon = ""
				currentNumber = ""
			}
		}
	}
	return cleansed
}
