package solution

import (
	_ "embed"
	"fmt"
	"reflect"
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
			name:            "with small sample",
			input:           sample0,
			expectedPartOne: 2028,
			expectedPartTwo: 1751,
		}, {
			name:            "with large sample",
			input:           sample1,
			expectedPartOne: 10092,
			expectedPartTwo: 9021,
		}, {
			name:            "with large input",
			input:           input,
			expectedPartOne: 1421727,
			expectedPartTwo: 1463160,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			grid, moves := parseInput(tst.input)
			if tst.expectedPartOne != -1 {
				if got := partOne(grid, moves, false); got != tst.expectedPartOne {
					t.Errorf("partOne() = %v, want %v", got, tst.expectedPartOne)
				}
			}

			if tst.expectedPartTwo != -1 {
				if got := partTwo(grid, moves, false); got != tst.expectedPartTwo {
					t.Errorf("partTwo() = %v, want %v", got, tst.expectedPartTwo)
				}
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	grid, moves := parseInput(sample0)
	expectedGrid := [][]byte{
		[]byte("########"),
		[]byte("#..O.O.#"),
		[]byte("##@.O..#"),
		[]byte("#...O..#"),
		[]byte("#.#.O..#"),
		[]byte("#...O..#"),
		[]byte("#......#"),
		[]byte("########"),
	}
	if !reflect.DeepEqual(expectedGrid, grid) {
		t.Errorf("parseInput()[0] = %v, want %v", grid, expectedGrid)
	}
	expectedMoves := "<^^>>>vv<v>>v<<"
	if expectedMoves != moves {
		t.Errorf("parseInput()[1] = %v, want %v", moves, expectedMoves)
	}
}

func parseInput(input string) ([][]byte, string) {
	lines := strings.Split(input, "\n")

	grid, i := [][]byte{}, 0
	for ; lines[i] != ""; i++ {
		grid = append(grid, []byte(lines[i]))
	}
	moves := strings.Join(lines[i:], "")

	return grid, moves
}

func findRobot(grid [][]byte) (int, int) {
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == '@' {
				return x, y
			}
		}
	}
	panic("cannot find robot")
}

func sumGPS(grid [][]byte) int {
	sum := 0
	h, w := len(grid), len(grid[0])
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 'O' || grid[y][x] == '[' {
				sum += 100*y + x
			}
		}
	}
	return sum
}

func canMove(x, y, dx, dy int, grid [][]byte) bool {
	cell := grid[y][x]
	if cell == '#' {
		return false
	}
	if cell == '.' {
		return true
	}
	nx, ny := x+dx, y+dy
	if dy != 0 && (cell == '[' || cell == ']') {
		if cell == ']' {
			nx -= 1
		}
		return canMove(nx, ny, dx, dy, grid) && canMove(nx+1, ny, dx, dy, grid)
	}
	return canMove(nx, ny, dx, dy, grid)
}

func move(x, y, dx, dy int, grid [][]byte) {
	cell := grid[y][x]
	if cell == '#' || cell == '.' {
		return
	}
	nx, ny := x+dx, y+dy
	if dy != 0 && (cell == '[' || cell == ']') {
		if cell == ']' {
			x, nx = x-1, nx-1
		}

		move(nx, ny, dx, dy, grid)
		grid[ny][nx] = grid[y][x]
		grid[y][x] = '.'

		move(nx+1, ny, dx, dy, grid)
		grid[ny][nx+1] = grid[y][x+1]
		grid[y][x+1] = '.'
	} else {
		move(nx, ny, dx, dy, grid)
		grid[ny][nx] = cell
		grid[y][x] = '.'
	}
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func deepCopy(src [][]byte) [][]byte {
	dst := make([][]byte, len(src))
	for i := range src {
		if src[i] != nil {
			dst[i] = make([]byte, len(src[i]))
			copy(dst[i], src[i])
		}
	}
	return dst
}

func partOne(grid [][]byte, moves string, debug bool) int {
	grid = deepCopy(grid)

	rx, ry := findRobot(grid)

	for _, m := range moves {
		var dx, dy int
		switch m {
		case '^':
			dx, dy = 0, -1
		case 'v':
			dx, dy = 0, 1
		case '>':
			dx, dy = 1, 0
		case '<':
			dx, dy = -1, 0
		}
		if canMove(rx, ry, dx, dy, grid) {
			move(rx, ry, dx, dy, grid)
			rx, ry = rx+dx, ry+dy
		}
		if debug {
			fmt.Println("Move:", string(m))
			printGrid(grid)
		}
	}

	return sumGPS(grid)
}

func doubleGrid(grid [][]byte) [][]byte {
	doubled := [][]byte{}

	for _, row := range grid {
		line := []byte{}

		for _, c := range row {
			switch c {
			case '#':
				line = append(line, "##"...)
			case 'O':
				line = append(line, "[]"...)
			case '.':
				line = append(line, ".."...)
			case '@':
				line = append(line, "@."...)
			}
		}

		doubled = append(doubled, line)
	}

	return doubled
}

func partTwo(grid [][]byte, moves string, debug bool) int {
	return partOne(doubleGrid(grid), moves, debug)
}
