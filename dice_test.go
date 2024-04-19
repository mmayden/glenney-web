package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetValue(t *testing.T) {
	// test cases
	testCases := []struct {
		newVal 		int
		dice		Dice
		expected	Dice
	}{
		{2, Dice{Value: 1, Size: 6}, Dice{Value: 1, Size: 2}},
		{11, Dice{Value: 3, Size: 6}, Dice{Value: 3, Size: 11}},
		{101, Dice{Value: 5, Size: 6}, Dice{Value: 5, Size: 101}},
	}

	// test case iteration
	for _, tc := range testCases {
		tc.dice.SetValue(tc.newVal)
		assert.Equal(t, tc.expected, tc.dice, "IntegerScan(%d) should return %d", tc.newVal, tc.expected)
	}
}