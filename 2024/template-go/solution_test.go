package solution

import (
	_ "embed"
	"reflect"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

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
			expectedPartOne: -1,
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: -1,
			expectedPartTwo: -1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			input := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(input); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(input); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	input := parseInput(sample)
	expected := []string{}
	if !reflect.DeepEqual(expected, input) {
		t.Errorf("parseInput() = %v, want %v", input, expected)
	}
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func partOne(input []string) int {
	return -1
}

func partTwo(input []string) int {
	return -1
}
