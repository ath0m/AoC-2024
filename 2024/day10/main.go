package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data = [][]int

func parse(filename string) *Data {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := Data{}
	for scanner.Scan() {
		row := []int{}
		for _, c := range scanner.Text() {
			row = append(row, int(c-'0'))
		}
		grid = append(grid, row)
	}

	if err = scanner.Err(); err != nil {
		panic("scanner error")
	}

	return &grid
}

type Pos struct {
	x, y int
}

func countTrailheads(start Pos, grid *Data) int {
	if (*grid)[start.x][start.y] != 0 {
		return 0
	}
	count := 0

	visited := map[Pos]bool{}
	q := []Pos{start}
	h, w := len(*grid), len((*grid)[0])

	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		_, ok := visited[v]
		if ok {
			continue
		}
		visited[v] = true

		num := (*grid)[v.x][v.y]
		if num == 9 {
			count += 1
			continue
		}

		for _, dv := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nv := Pos{v.x + dv.x, v.y + dv.y}
			if nv.x < 0 || nv.x >= h || nv.y < 0 || nv.y >= w {
				continue
			}

			if (*grid)[nv.x][nv.y] == num+1 {
				q = append(q, nv)
			}
		}
	}

	return count
}

func p1(grid *Data) int {
	result := 0
	for x := range *grid {
		for y := range (*grid)[x] {
			if (*grid)[x][y] == 0 {
				result += countTrailheads(Pos{x, y}, grid)
			}
		}
	}
	return result
}

func countTrailheads2(v Pos, grid *Data) int {
	count := 0
	h, w := len(*grid), len((*grid)[0])

	num := (*grid)[v.x][v.y]
	if num == 9 {
		return 1
	}

	for _, dv := range []Pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		nv := Pos{v.x + dv.x, v.y + dv.y}
		if nv.x < 0 || nv.x >= h || nv.y < 0 || nv.y >= w {
			continue
		}

		if (*grid)[nv.x][nv.y] == num+1 {
			count += countTrailheads2(nv, grid)
		}
	}

	return count
}

func p2(grid *Data) int {
	result := 0
	for x := range *grid {
		for y := range (*grid)[x] {
			if (*grid)[x][y] == 0 {
				result += countTrailheads2(Pos{x, y}, grid)
			}
		}
	}
	return result
}

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("argument is required")
	}
	data := parse("input.txt")
	switch args[1] {
	case "p1":
		fmt.Println("p1:", p1(data))
	case "p2":
		fmt.Println("p2:", p2(data))
	}
}
