package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompute(t *testing.T) {
	limits := map[Color]int{Red: 12, Green: 13, Blue: 14}
	result := compute("example.txt", limits)
	assert.EqualValues(t, 8, result)
}
