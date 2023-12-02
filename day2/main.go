package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Color string

const (
	Red   Color = "red"
	Blue        = "blue"
	Green       = "green"
)

func parseColor(input string) (Color, bool) {
	adapter := map[string]Color{"red": Red, "blue": Blue, "green": Green}
	c, ok := adapter[strings.ToLower(input)]
	return c, ok
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractGameId(input string) int {
	gameIDRegex := regexp.MustCompile(`Game (\d+)`)
	gameID := gameIDRegex.FindStringSubmatch(input)
	id, err := strconv.Atoi(gameID[1])
	check(err)
	return id
}

func extractDrawResult(input string) (int, Color) {
	drawResult := regexp.MustCompile(`(\d+) (red|blue|green)`)
	result := drawResult.FindStringSubmatch(input)
	n, err := strconv.Atoi(result[1])
	check(err)
	c, _ := parseColor(strings.Trim(result[2], " "))
	return n, c
}

func isRoundPossible(limits map[Color]int, draws []string) bool {
	roundPossible := true
	for _, c := range draws {
		n, color := extractDrawResult(c)
		limit, _ := limits[color]
		if n > limit {
			roundPossible = false
			break
		}
	}
	return roundPossible
}

func isGamePossible(limits map[Color]int, rounds []string) bool {
	gamePossible := true
	for _, r := range rounds {
		draws := strings.Split(r, ",")
		roundPossible := isRoundPossible(limits, draws)
		if !roundPossible {
			gamePossible = false
			break
		}
	}
	return gamePossible
}

func parseLine(line string) (string, []string) {
	parts := strings.Split(line, ":")
	return parts[0], strings.Split(parts[1], ";")
}

func computePossibleGames(filename string, limits map[Color]int) int {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		game, rounds := parseLine(line)
		gamePossible := isGamePossible(limits, rounds)

		gameID := extractGameId(game)
		if gamePossible {
			result += gameID
		}

	}

	return result
}

func computePower(filename string) int {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	var result int
	for scanner.Scan() {
		line := scanner.Text()
		_, rounds := parseLine(line)
		drawsPerColor := map[Color][]int{Red: {}, Green: {}, Blue: {}}
		for _, r := range rounds {
			draws := strings.Split(r, ",")
			for _, c := range draws {
				n, color := extractDrawResult(c)
				drawsPerColor[color] = append(drawsPerColor[color], n)
			}
		}
		power := 1
		fmt.Println(drawsPerColor)
		for _, draws := range drawsPerColor {
			power = power * slices.Max(draws)
		}
		result += power
		fmt.Println(result, power)
	}

	return result
}
func main() {
	limits := map[Color]int{Red: 12, Green: 13, Blue: 14}
	result := computePossibleGames("day2/input.txt", limits)
	power := computePower("day2/input.txt")
	fmt.Println("result", result, power)
}
