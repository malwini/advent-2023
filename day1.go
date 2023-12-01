package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func compute() int {
	file, err := os.Open("day1_input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		var numbersInLine []int
		for i, b := range line {
			if unicode.IsNumber(b) {
				n, err := strconv.Atoi(line[i : i+1])
				check(err)
				numbersInLine = append(numbersInLine, n)
			}
		}
		result += numbersInLine[0]*10 + numbersInLine[len(numbersInLine)-1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	result := compute()
	fmt.Println(result)

}
