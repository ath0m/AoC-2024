package solution

import (
	_ "embed"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

//go:embed sample1.txt
var sample1 string

//go:embed sample2.txt
var sample2 string

func TestSolution(t *testing.T) {
	type test struct {
		name            string
		input           string
		expectedPartOne string
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			expectedPartOne: "4,6,3,5,6,3,5,2,1,0",
			expectedPartTwo: -1,
		}, {
			name:            "with second sample",
			input:           sample1,
			expectedPartOne: "0,3,5,4,3,0",
			expectedPartTwo: -1,
		}, {
			name:            "with third sample",
			input:           sample2,
			expectedPartOne: "2,4,1,1,7,5,1,5,4,3,0,3,5,5,3,0",
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: "7,6,1,5,3,1,4,2,6",
			expectedPartTwo: 164541017976509,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			input := parseInput(tst.input)
			if tst.expectedPartOne != "" {
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

type State struct {
	regA, regB, regC int
	program          []int
}

func TestParseInput(t *testing.T) {
	input := parseInput(sample)
	expected := State{
		729, 0, 0,
		[]int{0, 1, 5, 4, 3, 0},
	}
	if !reflect.DeepEqual(expected, input) {
		t.Errorf("parseInput() = %v, want %v", input, expected)
	}
}

func parseInput(input string) State {
	lines := strings.Split(input, "\n")

	regA, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	regB, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	regC, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	nums := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	program := []int{}
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		program = append(program, val)
	}

	return State{
		regA, regB, regC,
		program,
	}
}

type Computer struct {
	regA, regB, regC int
	program          []int

	pointer int
	outputs []int
}

func (c *Computer) load(state State) {
	c.regA = state.regA
	c.regB = state.regB
	c.regC = state.regC

	c.program = make([]int, len(state.program))
	copy(c.program, state.program)

	c.pointer = 0
}

func (c Computer) combo(operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	}
	panic("unsupported operand")
}

func (c *Computer) cycle() {
	opcode, operand := c.program[c.pointer], c.program[c.pointer+1]
	switch opcode {
	case 0: // adv
		c.regA = c.regA >> c.combo(operand)
	case 1: // bxl
		c.regB = c.regB ^ operand
	case 2: // bst
		c.regB = c.combo(operand) & 7
	case 3: // jnz
		if c.regA != 0 {
			c.pointer = operand
			return
		}
	case 4: // bxc
		c.regB = c.regB ^ c.regC
	case 5: // out
		c.outputs = append(c.outputs, c.combo(operand)&7)
	case 6: // bdv
		c.regB = c.regA >> c.combo(operand)
	case 7: // cdv
		c.regC = c.regA >> c.combo(operand)
	}
	c.pointer += 2
}

func (c *Computer) run(limit int) {
	cnt := 0
	for c.pointer < len(c.program) && (cnt < limit || limit < 0) {
		c.cycle()
		cnt += 1
	}
}

func vecToString(vec []int) string {
	output := []string{}
	for _, out := range vec {
		output = append(output, strconv.Itoa(out))
	}
	return strings.Join(output, ",")
}

func partOne(state State) string {
	comp := Computer{}
	comp.load(state)
	comp.run(-1)
	return vecToString(comp.outputs)
}

func findRegA(a int, seq []int, state State) (int, bool) {
	n := len(seq)
	if n == 0 {
		return a, true
	}

	for i := 0; i < 8; i++ {
		regA := (a << 3) | i
		comp := Computer{}
		comp.load(state)
		comp.regA = regA
		comp.run(-1)
		if comp.outputs[0] == seq[n-1] {
			res, ok := findRegA(regA, seq[:n-1], state)
			if ok {
				return res, true
			}
		}
	}

	return 0, false
}

func partTwo(state State) int {
	regA, ok := findRegA(0, state.program, state)
	if ok {
		return regA
	}
	return -1
}
