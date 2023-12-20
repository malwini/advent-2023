package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputeSum(t *testing.T) {
	result := computeSum("example.txt")
	assert.EqualValues(t, 13, result)
}

func TestComputeSumWithAdditionalCards(t *testing.T) {
	result := computeSumWithAdditionalCards("example.txt")
	assert.EqualValues(t, 30, result)
}
