package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	got := compute()
	assert.EqualValues(t, 142, got)
}
