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

	lines := ""
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		lines += line
	}

	lines = strings.ReplaceAll(lines, "\n", "-")
	lines = strings.ReplaceAll(lines, "do()", "🙃")
	lines = strings.ReplaceAll(lines, "don't()", "😎")
	clean := cleanse2(lines)
	value := processCleansed2(clean)
	fmt.Printf("%v\n", value)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processCleansed2(clean string) int {
	value := 0
	split := strings.Fields(clean)
	enabled := true
	for _, compute := range split {
		// fmt.Println(compute)
		if compute == "yes" {
			enabled = true
		} else if compute == "no" {
			enabled = false
		} else {
			if enabled {
				computeNaked := strings.Replace(compute, "mul(", "", -1)
				computeNaked = strings.Replace(computeNaked, ")", "", -1)
				computeSplit := strings.Split(computeNaked, ",")
				if first, err := strconv.Atoi(computeSplit[0]); err == nil {
					if second, err := strconv.Atoi(computeSplit[1]); err == nil {
						value += (first * second)
					}
				}
			}
		}
	}
	return value
}

func cleanse2(input string) string {
	cleansed := ""
	currentLexicon := ""
	currentNumber := ""
	for _, tokenRune := range input {
		token := string(tokenRune)
		if token == "🙃" {
			fmt.Println("DO()")
			cleansed += "yes"
			cleansed += " "
		}
		if token == "😎" {
			fmt.Println("DONT()")
			cleansed += "no"
			cleansed += " "
		}
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
