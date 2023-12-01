package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestA(t *testing.T) {
	got := compute("day1_a_example.txt")
	assert.EqualValues(t, 142, got)
}

func TestB(t *testing.T) {
	got := compute("day1_b_example.txt")
	assert.EqualValues(t, 281, got)
}
