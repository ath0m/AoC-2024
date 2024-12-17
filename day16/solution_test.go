package solution

import (
	"container/heap"
	_ "embed"
	"fmt"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

//go:embed sample0.txt
var sample0 string

//go:embed sample1.txt
var sample1 string

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
			input:           sample0,
			expectedPartOne: 7036,
			expectedPartTwo: 45,
		}, {
			name:            "with second sample",
			input:           sample1,
			expectedPartOne: 11048,
			expectedPartTwo: 64,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 109516,
			expectedPartTwo: 568,
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

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines
}

func findCell(maze []string, cell byte) Pos {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == cell {
				return Pos{Vec{x, y}, East}
			}
		}
	}
	panic("cannot find reindeer")
}

func mod(a, b int) int {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}

func (p Pos) move(d Dir) (Pos, int) {
	var npos Pos
	switch d {
	case East:
		npos = Pos{Vec{p.xy.x + 1, p.xy.y}, d}
	case North:
		npos = Pos{Vec{p.xy.x, p.xy.y - 1}, d}
	case West:
		npos = Pos{Vec{p.xy.x - 1, p.xy.y}, d}
	case South:
		npos = Pos{Vec{p.xy.x, p.xy.y + 1}, d}
	}
	turns := min(mod(int(p.dir)-int(d), 4), mod(int(d)-int(p.dir), 4))
	return npos, 1 + 1000*turns
}

func partOne(maze []string) int {
	start := findCell(maze, 'S')

	pq := PriorityQueue{&Item{
		start, start, 0, 0,
	}}
	heap.Init(&pq)

	visited := map[Vec]bool{}

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		_, ok := visited[item.value.xy]
		if ok {
			continue
		}
		visited[item.value.xy] = true

		if maze[item.value.xy.y][item.value.xy.x] == 'E' {
			return item.priority
		}

		for _, d := range []Dir{East, North, West, South} {
			npos, cost := item.value.move(d)
			if maze[npos.xy.y][npos.xy.x] != '#' {
				heap.Push(&pq, &Item{
					value:    npos,
					prev:     item.value,
					priority: item.priority + cost,
				})
			}
		}
	}

	panic("cannot find end")
}

func printCells(maze []string, path map[Vec]bool) {
	for y := range maze {
		for x := range maze[y] {
			_, ok := path[Vec{x, y}]
			if ok {
				fmt.Print("O")
			} else {
				fmt.Print(string(maze[y][x]))
			}
		}
		fmt.Println("")
	}
}

func partTwo(maze []string) int {
	start := findCell(maze, 'S')

	pq := PriorityQueue{&Item{
		start, start, 0, 0,
	}}
	heap.Init(&pq)

	dist := map[Pos]int{}
	prevs := map[Pos][]Pos{}

	total := -1

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		if total >= 0 && item.priority > total {
			continue
		}

		d, ok := dist[item.value]
		if !ok || item.priority < d {
			dist[item.value] = item.priority
			prevs[item.value] = []Pos{item.prev}
		} else if d == item.priority {
			prevs[item.value] = append(prevs[item.value], item.prev)
		} else {
			continue
		}

		if maze[item.value.xy.y][item.value.xy.x] == 'E' {
			total = item.priority
			continue
		}

		// move forward
		var np Vec
		switch item.value.dir {
		case East:
			np = Vec{item.value.xy.x + 1, item.value.xy.y}
		case North:
			np = Vec{item.value.xy.x, item.value.xy.y - 1}
		case West:
			np = Vec{item.value.xy.x - 1, item.value.xy.y}
		case South:
			np = Vec{item.value.xy.x, item.value.xy.y + 1}
		}
		if maze[np.y][np.x] != '#' {
			heap.Push(&pq, &Item{
				value: Pos{
					np,
					item.value.dir,
				},
				prev:     item.value,
				priority: item.priority + 1,
			})
		}

		// turn 90, -90 degree
		for _, dd := range []int{-1, 1} {
			ndir := Dir(mod(int(item.value.dir)+dd, 4))
			heap.Push(&pq, &Item{
				value: Pos{
					item.value.xy,
					ndir,
				},
				prev:     item.value,
				priority: item.priority + 1000,
			})
		}
	}

	end := findCell(maze, 'E')
	cells := map[Vec]bool{}
	visited := map[Pos]bool{}

	q := []Pos{{end.xy, East}, {end.xy, North}, {end.xy, West}, {end.xy, South}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		_, ok := visited[v]
		if ok {
			continue
		}
		visited[v] = true
		cells[v.xy] = true

		for _, u := range prevs[v] {
			if u != v {
				q = append(q, u)
			}
		}
	}

	// printCells(maze, cells)

	return len(cells)
}
