package main

import (
	"bufio"
	"fmt"
	"os"
)

func parse(filename string) *[]string {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic("scanner error")
	}

	return &lines
}

type Point struct {
	x, y int
}

func (p Point) sub(other Point) Vec {
	return Vec{p.x - other.x, p.y - other.y}
}

func (p Point) add(vec Vec) Point {
	return Point{p.x + vec.x, p.y + vec.y}
}

type Vec Point

func (vec Vec) scale(alpha int) Vec {
	return Vec{vec.x * alpha, vec.y * alpha}
}

func extractAntennas(grid *[]string) *map[rune]([]Point) {
	antennas := map[rune]([]Point){}
	for i, row := range *grid {
		for j, ch := range row {
			if ch != '.' {
				p := Point{i, j}
				antennas[ch] = append(antennas[ch], p)
			}
		}
	}
	return &antennas
}

func p1(grid *[]string) int {
	antennas := extractAntennas(grid)
	antinodes := map[Point]bool{}

	h, w := len(*grid), len((*grid)[0])

	for _, points := range *antennas {
		for i := 0; i < len(points); i++ {
			p1 := points[i]
			for j := 0; j < len(points); j++ {
				if i == j {
					continue
				}
				p2 := points[j]

				antinode := p1.add(p2.sub(p1).scale(2))
				x, y := antinode.x, antinode.y
				if 0 <= x && x < h && 0 <= y && y < w {
					antinodes[antinode] = true
				}
			}
		}
	}

	return len(antinodes)
}

func p2(grid *[]string) int {
	antennas := extractAntennas(grid)
	antinodes := map[Point]bool{}

	h, w := len(*grid), len((*grid)[0])

	for _, points := range *antennas {
		for i := 0; i < len(points); i++ {
			p1 := points[i]
			for j := 0; j < len(points); j++ {
				if i == j {
					continue
				}
				p2 := points[j]

				alpha := 1
				for true {
					antinode := p1.add(p2.sub(p1).scale(alpha))
					x, y := antinode.x, antinode.y
					if x < 0 || x >= h || y < 0 || y >= w {
						break
					}
					antinodes[antinode] = true
					alpha += 1
				}

			}
		}
	}

	return len(antinodes)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("argument is required")
	}
	grid := parse("input.txt")
	switch args[1] {
	case "p1":
		fmt.Println("p1:", p1(grid))
	case "p2":
		fmt.Println("p2:", p2(grid))
	}
}
