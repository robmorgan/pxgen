package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var TestChars = []byte("AAA")

func TestLengthParam(t *testing.T) {
	assert.Equal(t, 12, len(rand_char(12, StdChars)), "length should be 12")
}

func TestCharsParam(t *testing.T) {
	assert.Equal(t, "A", rand_char(1, TestChars), "length should be 12")
}
