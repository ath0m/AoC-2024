package solution

import (
	_ "embed"
	"reflect"
	"regexp"
	"strconv"
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
		width           int
		height          int
		expectedPartOne int
		expectedPartTwo int
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			width:           11,
			height:          7,
			expectedPartOne: 12,
			expectedPartTwo: -1,
		}, {
			name:            "with large input",
			input:           input,
			width:           101,
			height:          103,
			expectedPartOne: 231852216,
			expectedPartTwo: 8159,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			robots := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(robots, tst.width, tst.height); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(robots, tst.width, tst.height); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	robots := parseInput(sample)
	expected := []Robot{
		{pos: Vec{0, 4}, vel: Vec{3, -3}},
		{pos: Vec{6, 3}, vel: Vec{-1, -3}},
		{pos: Vec{10, 3}, vel: Vec{-1, 2}},
		{pos: Vec{2, 0}, vel: Vec{2, -1}},
		{pos: Vec{0, 0}, vel: Vec{1, 3}},
		{pos: Vec{3, 0}, vel: Vec{-2, -2}},
		{pos: Vec{7, 6}, vel: Vec{-1, -3}},
		{pos: Vec{3, 0}, vel: Vec{-1, -2}},
		{pos: Vec{9, 3}, vel: Vec{2, 3}},
		{pos: Vec{7, 3}, vel: Vec{-1, 2}},
		{pos: Vec{2, 4}, vel: Vec{2, -3}},
		{pos: Vec{9, 5}, vel: Vec{-3, -3}},
	}
	if !reflect.DeepEqual(expected, robots) {
		t.Errorf("parseInput() = %v, want %v", robots, expected)
	}
}

func parseInput(input string) []Robot {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`p=([-]?\d+),([-]?\d+) v=([-]?\d+),([-]?\d+)`)

	robots := []Robot{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		if len(matches) == 5 {
			px, _ := strconv.Atoi(matches[1])
			py, _ := strconv.Atoi(matches[2])
			vx, _ := strconv.Atoi(matches[3])
			vy, _ := strconv.Atoi(matches[4])

			robot := Robot{
				pos: Vec{px, py},
				vel: Vec{vx, vy},
			}
			robots = append(robots, robot)
		}
	}

	return robots
}

type Vec struct {
	x, y int
}

func (vec Vec) scale(alpha int) Vec {
	return Vec{vec.x * alpha, vec.y * alpha}
}

func (vec Vec) add(other Vec) Vec {
	return Vec{vec.x + other.x, vec.y + other.y}
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}

func (vec Vec) mod(other Vec) Vec {
	return Vec{mod(vec.x, other.x), mod(vec.y, other.y)}
}

type Robot struct {
	pos Vec
	vel Vec
}

func (robot Robot) simulate(steps int, board Vec) Vec {
	return robot.pos.add(robot.vel.scale(steps)).mod(board)
}

type Quadrant struct {
	x0, y0, x1, y1 int
}

func (q Quadrant) contain(vec Vec) bool {
	return q.x0 <= vec.x && vec.x <= q.x1 && q.y0 <= vec.y && vec.y <= q.y1
}

func partOne(robots []Robot, width, height int) int {
	steps := 100
	size := Vec{width, height}

	q0 := Quadrant{0, 0, width/2 - 1, height/2 - 1}
	q1 := Quadrant{width/2 + 1, 0, width - 1, height/2 - 1}
	q2 := Quadrant{0, height/2 + 1, width/2 - 1, height - 1}
	q3 := Quadrant{width/2 + 1, height/2 + 1, width - 1, height - 1}

	c0, c1, c2, c3 := 0, 0, 0, 0

	for _, robot := range robots {
		pos := robot.simulate(steps, size)

		switch {
		case q0.contain(pos):
			c0 += 1
		case q1.contain(pos):
			c1 += 1
		case q2.contain(pos):
			c2 += 1
		case q3.contain(pos):
			c3 += 1
		}
	}
	return c0 * c1 * c2 * c3
}

func partTwo(robots []Robot, width, height int) int {
	size := Vec{width, height}

	q0 := Quadrant{0, 0, width/2 - 1, height/2 - 1}
	q1 := Quadrant{width/2 + 1, 0, width - 1, height/2 - 1}
	q2 := Quadrant{0, height/2 + 1, width/2 - 1, height - 1}
	q3 := Quadrant{width/2 + 1, height/2 + 1, width - 1, height - 1}

	minFactor := 1000000000
	var minStep int

	for step := 1; step < width*height; step++ {
		c0, c1, c2, c3 := 0, 0, 0, 0

		for _, robot := range robots {
			pos := robot.simulate(step, size)

			switch {
			case q0.contain(pos):
				c0 += 1
			case q1.contain(pos):
				c1 += 1
			case q2.contain(pos):
				c2 += 1
			case q3.contain(pos):
				c3 += 1
			}
		}

		factor := c0 * c1 * c2 * c3
		if factor < minFactor {
			minFactor = factor
			minStep = step
		}
	}

	return minStep
}
