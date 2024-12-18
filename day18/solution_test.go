package solution

import (
	"container/heap"
	_ "embed"
	"fmt"
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
		steps           int
		expectedPartOne int
		expectedPartTwo string
	}

	tests := []test{
		{
			name:            "with sample",
			input:           sample,
			width:           7,
			height:          7,
			steps:           12,
			expectedPartOne: 22,
			expectedPartTwo: "6,1",
		}, {
			name:            "with large input",
			input:           input,
			width:           71,
			height:          71,
			steps:           1024,
			expectedPartOne: 290,
			expectedPartTwo: "64,54",
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			input := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(input, tst.width, tst.height, tst.steps); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != "" {
				if got := partTwo(input, tst.width, tst.height); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})
	}
}

func parseInput(input string) []Vec {
	lines := strings.Split(input, "\n")
	points := []Vec{}

	for _, line := range lines {
		if len(line) > 0 {
			xy := strings.Split(line, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			points = append(points, Vec{x, y})
		}
	}

	return points
}

func dijkstra(start, end Vec, grid [][]int) (int, bool) {
	pq := PriorityQueue[Vec]{&Item[Vec]{start, 0}}
	heap.Init(&pq)

	h, w := len(grid), len(grid[0])

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item[Vec])
		p := item.value
		if p == end {
			return item.priority, true
		}
		if grid[p.y][p.x] == 1 {
			continue
		}
		grid[p.y][p.x] = 1

		for _, dp := range []Vec{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			np := Vec{p.x + dp.x, p.y + dp.y}
			if 0 <= np.x && np.x < w && 0 <= np.y && np.y < h && grid[np.y][np.x] == 0 {
				heap.Push(&pq, &Item[Vec]{
					np, item.priority + 1,
				})
			}
		}
	}

	return -1, false
}

func partOne(points []Vec, width, height int, steps int) int {
	grid := make([][]int, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)
	}
	for i := 0; i < steps; i++ {
		p := points[i]
		grid[p.y][p.x] = 1
	}

	start := Vec{0, 0}
	end := Vec{width - 1, height - 1}

	dist, ok := dijkstra(start, end, grid)
	if ok {
		return dist
	}

	return -1
}

func partTwo(points []Vec, width, height int) string {
	l, r := 0, len(points)
	for l < r {
		m := (l + r) / 2
		if partOne(points, width, height, m) != -1 {
			l = m + 1
		} else {
			r = m
		}
	}
	p := points[l-1]
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
