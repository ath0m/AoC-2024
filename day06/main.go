package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse(filename string) *[][]byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]byte{}
	for scanner.Scan() {
		raw := scanner.Bytes()
		line := make([]byte, len(raw))
		copy(line, raw)
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scanning file: %v", err)
	}
	return &grid
}

func findGuard(grid *[][]byte) (int, int, int, int) {
	h, w := len(*grid), len((*grid)[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cell := (*grid)[i][j]
			switch cell {
			case '^':
				return i, j, -1, 0
			case 'v':
				return i, j, 1, 0
			case '<':
				return i, j, 0, -1
			case '>':
				return i, j, 0, 1
			}
		}
	}
	panic("cannot find guard")
}

func p1(grid *[][]byte) int {
	cnt := 0

	h, w := len(*grid), len((*grid)[0])
	gi, gj, gdi, gdj := findGuard(grid)

	for true {
		if (*grid)[gi][gj] != 'X' {
			(*grid)[gi][gj] = 'X'
			cnt += 1
		}
		ngi, ngj := gi+gdi, gj+gdj
		if ngi < 0 || ngi >= h || ngj < 0 || ngj >= w {
			// guard is leaving the grid
			break
		}

		if (*grid)[ngi][ngj] == '#' {
			// rotate guard 90' clockwise
			gdi, gdj = gdj, -gdi
			continue
		}

		gi, gj = ngi, ngj
	}

	return cnt
}

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

func rot90(dir Dir) Dir {
	switch dir {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	}
	panic("Unsupproted direction")
}

func move(i, j int, dir Dir) (int, int) {
	switch dir {
	case UP:
		return i - 1, j
	case RIGHT:
		return i, j + 1
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	}
	panic("Unsupproted direction")
}

type Pos struct {
	i, j int
	dir  Dir
}

func (p Pos) move() Pos {
	ni, nj := move(p.i, p.j, p.dir)
	return Pos{ni, nj, p.dir}
}

func (p Pos) valid(h, w int) bool {
	return 0 <= p.i && p.i < h && 0 <= p.j && p.j < w
}

func (p Pos) rot90() Pos {
	return Pos{p.i, p.j, rot90(p.dir)}
}

func printGrid(grid *[][]byte) {
	for _, row := range *grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func simulateLoop(grid *[][]byte, gi, gj int, dir Dir, oi, oj int) bool {
	tmp := (*grid)[oi][oj]
	(*grid)[oi][oj] = 'O'
	defer func() {
		(*grid)[oi][oj] = tmp
	}()
	p := Pos{gi, gj, dir}
	h, w := len(*grid), len((*grid)[0])

	visited := map[Pos]bool{}

	for true {
		_, ok := visited[p]

		if ok {
			// printGrid(grid)
			return true
		}
		visited[p] = true

		np := p.move()
		if !np.valid(h, w) {
			// guard is leaving the grid
			break
		}

		if (*grid)[np.i][np.j] == '#' || (*grid)[np.i][np.j] == 'O' {
			// rotate guard 90' clockwise
			p = p.rot90()
			continue
		}

		p = np
	}
	return false
}

type Pair struct {
	x, y int
}

func p2(grid *[][]byte) int {
	cnt := 0
	h, w := len(*grid), len((*grid)[0])
	gsi, gsj, _, _ := findGuard(grid)
	sdir := UP

	gi, gj, dir := gsi, gsj, sdir
	checked := map[Pair]bool{}

	for true {
		ngi, ngj := move(gi, gj, dir)
		if ngi < 0 || ngi >= h || ngj < 0 || ngj >= w {
			// guard is leaving the grid
			break
		}

		if (*grid)[ngi][ngj] == '#' {
			// rotate guard 90' clockwise
			dir = rot90(dir)
			continue
		}

		obs := Pair{ngi, ngj}
		_, ok := checked[obs]
		if !ok {
			if simulateLoop(grid, gsi, gsj, sdir, obs.x, obs.y) {
				cnt += 1
			}
			checked[obs] = true
		}

		gi, gj = ngi, ngj
	}

	return cnt
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatalf("expects only one argument to be provided: p1, p2")
	}
	grid := parse("input.txt")
	switch args[1] {
	case "p1":
		fmt.Println("p1:", p1(grid))
	case "p2":
		fmt.Println("p2:", p2(grid))
	}
}
