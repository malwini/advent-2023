package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompute(t *testing.T) {
	result := computeSum("example.txt")
	assert.EqualValues(t, 4361, result)
}

func TestIsSymbol(t *testing.T) {
	assert.True(t, isSymbol('+'))
	assert.False(t, isSymbol('.'))
	assert.False(t, isSymbol('9'))
}
