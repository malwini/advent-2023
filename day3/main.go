package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type PositionRange struct {
	rowNumber         int
	FirstColumnNumber int
	LastColumnNumber  int
}

type Position struct {
	row    int
	column int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(fileName string) map[int]string {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	lines := map[int]string{}
	var rowNumber int
	for scanner.Scan() {
		lines[rowNumber] = scanner.Text()
		rowNumber++
	}
	return lines
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func hasAdjacentSymbol(positions PositionRange, lines map[int]string) bool {
	for p := positions.FirstColumnNumber; p <= positions.LastColumnNumber; p++ {
		for i := p - 1; i < p+1; i++ {
			for j := positions.rowNumber - 1; j <= positions.rowNumber+1; j++ {
				line, ok := lines[j]
				if j < 0 || !ok {
					continue
				}
				if i < 0 || i >= len(line) { // if i is out of range, skip
					continue
				}
				adjacent := rune(line[i])
				if isSymbol(adjacent) {
					return true
				}
			}
		}
	}
	return false
}

func hasAdjacentStar(positions PositionRange, lines map[int]string) (Position, bool) {
	for p := positions.FirstColumnNumber; p <= positions.LastColumnNumber; p++ {
		for i := p - 1; i < p+1; i++ {
			for j := positions.rowNumber - 1; j <= positions.rowNumber+1; j++ {
				line, ok := lines[j]
				if j < 0 || !ok {
					continue
				}
				if i < 0 || i >= len(line) { // if i is out of range, skip
					continue
				}
				adjacent := rune(line[i])
				if adjacent == '*' {
					return Position{j, i}, true
				}
			}
		}
	}
	return Position{}, false
}

func computeSum(fileName string) int {
	lines := readLines(fileName)
	var result int
	for rowNumber, line := range lines {
		regexNumbersInLine := regexp.MustCompile(`\d+`)
		numbersInLine := regexNumbersInLine.FindAllStringIndex(line, -1)
		for _, n := range numbersInLine {
			ok := hasAdjacentSymbol(PositionRange{rowNumber, n[0], n[1]}, lines)
			if ok {
				number, _ := strconv.Atoi(line[n[0]:n[1]])
				result += number
			}
		}
	}
	return result
}

func computeGearRatio(fileName string) int {
	lines := readLines(fileName)
	var result int
	numbers := map[Position]int{}
	for rowNumber, line := range lines {
		regexNumbersInLine := regexp.MustCompile(`\d+`)
		numbersInLine := regexNumbersInLine.FindAllStringIndex(line, -1)
		for _, n := range numbersInLine {
			p, hasAdjStart := hasAdjacentStar(PositionRange{rowNumber, n[0], n[1]}, lines)
			if hasAdjStart {
				existingNumber, ok := numbers[p]
				number, _ := strconv.Atoi(line[n[0]:n[1]])
				if ok {
					result += existingNumber * number
				} else {
					numbers[p] = number
				}
			}
		}
	}
	return result
}

func main() {
	result := computeSum("day3/input.txt")
	fmt.Println("result", result)

	r := computeGearRatio("day3/input.txt")
	fmt.Println("gear", r)
}
