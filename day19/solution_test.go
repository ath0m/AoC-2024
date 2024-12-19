package main

import (
	_ "embed"
	"testing"
)

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: 6,
			expectedPartTwo: 16,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 304,
			expectedPartTwo: 705756472327497,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			patterns, designs := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(patterns, designs); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(patterns, designs); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})
	}
}
