package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Card struct {
	Id             int
	OwnNumbers     []int
	WinningNumbers []int
}

func (receiver Card) NumberOfWins() int {
	numberOfWins := 0
	for _, n := range receiver.OwnNumbers {
		if slices.Contains(receiver.WinningNumbers, n) {
			numberOfWins++
		}
	}
	return numberOfWins
}
func (receiver Card) Score() int {
	return int(math.Pow(2.0, float64(receiver.NumberOfWins())-1))
}

func stringListToIntList(input []string) []int {
	var result []int
	for _, s := range input {
		n, err := strconv.Atoi(s)
		check(err)
		result = append(result, n)
	}
	return result
}

func parseNumberList(input string) []int {
	numbers := regexp.MustCompile(`[0-9]+`).FindAllString(input, -1)
	intNumbers := stringListToIntList(numbers)
	return intNumbers
}
func parseCardId(input string) int {
	idRegex := regexp.MustCompile(`Card\s+([0-9]+)`)
	idGroups := idRegex.FindStringSubmatch(input)
	id, err := strconv.Atoi(idGroups[1])
	check(err)
	return id
}

func parseCard(input string) Card {
	cardRegex := regexp.MustCompile(`(Card .*): (.*) \| (.*)`)
	cardGroups := cardRegex.FindStringSubmatch(input)
	ownNumbers := parseNumberList(cardGroups[2])
	winningNumbers := parseNumberList(cardGroups[3])
	return Card{parseCardId(cardGroups[1]), ownNumbers, winningNumbers}
}

func computeSum(fileName string) int {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		result += card.Score()
	}
	return result
}

func computeSumWithAdditionalCards(fileName string) int {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	var result int
	originalCards := []Card{}
	for scanner.Scan() {
		line := scanner.Text()
		card := parseCard(line)
		originalCards = append(originalCards, card)
	}
	final := map[int]int{}

	for p := 1; p <= len(originalCards); p++ {
		card := originalCards[p-1]
		pos := card.Id
		final[pos] += 1
		bound := min(len(originalCards)+1, pos+card.NumberOfWins()+1)
		for i := pos + 1; i < bound; i++ {
			final[i] += final[pos]
		}
		result += final[pos]
	}
	return result
}

func main() {
	fmt.Println("result", computeSum("day4/input.txt"))

	fmt.Println("result with additional", computeSumWithAdditionalCards("day4/input.txt"))
}
