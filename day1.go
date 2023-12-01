package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseWordToDigit(input string) (int, bool) {
	digitsAsWords := map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9}
	r, exists := digitsAsWords[strings.ToLower(input)]
	if !exists {
		return -1, false
	}
	return r, true
}

func compute(filename string) int {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		var numbersInLine []int
		for i, b := range line {
			if unicode.IsNumber(b) {
				n, err := strconv.Atoi(line[i : i+1])
				check(err)
				numbersInLine = append(numbersInLine, n)
			}
			for j := i; j <= len(line); j++ {
				n, isDigit := parseWordToDigit(line[i:j])
				if isDigit {
					numbersInLine = append(numbersInLine, n)
				}
			}
		}
		result += numbersInLine[0]*10 + numbersInLine[len(numbersInLine)-1]
		fmt.Println(numbersInLine[0]*10 + numbersInLine[len(numbersInLine)-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	result := compute("day1_input.txt")
	fmt.Println(result)

}
