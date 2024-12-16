package main

import (
	"bufio"
	"fmt"
	"os"
)

type Data = []string

func parse(filename string) *Data {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := Data{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic("scanner error")
	}

	return &lines
}

type Pair struct {
	x, y int
}

func countRegion(grid *Data, x, y int, region byte, visited *[][]bool) (int, int) {
	h, w := len(*grid), len((*grid)[0])
	cell := (*grid)[x][y]

	(*visited)[x][y] = true

	area := 1
	perimeter := 0

	for _, d := range []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx, ny := x+d.x, y+d.y
		if nx < 0 || nx >= h || ny < 0 || ny >= w || (*grid)[nx][ny] != cell {
			perimeter += 1
		} else if !(*visited)[nx][ny] {
			a, p := countRegion(grid, nx, ny, region, visited)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func p1(data *Data) int {
	h, w := len(*data), len((*data)[0])
	visited := [][]bool{}
	for i := 0; i < h; i++ {
		visited = append(visited, make([]bool, w))
	}

	result := 0
	for x := range *data {
		for y := range (*data)[x] {
			if !visited[x][y] {
				cell := (*data)[x][y]
				a, p := countRegion(data, x, y, cell, &visited)
				result += a * p
			}
		}
	}
	return result
}

type Segment struct {
	x, y   int
	dx, dy int
}

func countAreaAndSegments(
	grid Data,
	x, y int,
	region byte,
	visited [][]bool,
	segments map[Segment]bool,
) int {
	h, w := len(grid), len(grid[0])
	cell := grid[x][y]

	visited[x][y] = true

	area := 1

	for _, d := range []Pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx, ny := x+d.x, y+d.y
		if nx < 0 || nx >= h || ny < 0 || ny >= w || grid[nx][ny] != cell {
			var segment Segment
			switch d {
			case Pair{-1, 0}:
				segment = Segment{x, y, 0, 1}
			case Pair{1, 0}:
				segment = Segment{x + 1, y, 0, 1}
			case Pair{0, -1}:
				segment = Segment{x, y, 1, 0}
			case Pair{0, 1}:
				segment = Segment{x, y + 1, 1, 0}
			}
			segments[segment] = true
		} else if !visited[nx][ny] {
			a := countAreaAndSegments(grid, nx, ny, region, visited, segments)
			area += a
		}
	}

	return area
}

func getAnySegment(m map[Segment]bool) (Segment, bool) {
	for key := range m {
		return key, true // Return the first key found
	}
	return Segment{}, false // If the map is empty, indicate that no key was found
}

func countSides(segments map[Segment]bool) int {
	sides := 0
	for len(segments) > 0 {
		segment, _ := getAnySegment(segments)
		x := segment.x
		y := segment.y
		dx := segment.dx
		dy := segment.dy

		sides += 1
		delete(segments, segment)

		for i := 1; ; i++ {
			seg := Segment{x + dx*i, y + dy*i, dx, dy}
			_, ok := segments[seg]
			if !ok {
				break
			}
			delete(segments, seg)
		}

		for i := 1; ; i++ {
			seg := Segment{x - dx*i, y - dy*i, dx, dy}
			_, ok := segments[seg]
			if !ok {
				break
			}
			delete(segments, seg)
		}
	}
	return sides
}

func p2(data *Data) int {
	h, w := len(*data), len((*data)[0])
	visited := [][]bool{}
	for i := 0; i < h; i++ {
		visited = append(visited, make([]bool, w))
	}

	result := 0
	for x := range *data {
		for y := range (*data)[x] {
			if !visited[x][y] {
				cell := (*data)[x][y]
				segments := map[Segment]bool{}
				a := countAreaAndSegments(*data, x, y, cell, visited, segments)
				result += a * countSides(segments)
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
