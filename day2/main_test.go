package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputePossibleGames(t *testing.T) {
	limits := map[Color]int{Red: 12, Green: 13, Blue: 14}
	result := computePossibleGames("example.txt", limits)
	assert.EqualValues(t, 8, result)
}

func TestComputePower(t *testing.T) {
	result := computePower("example.txt")
	assert.EqualValues(t, 2286, result)
}
