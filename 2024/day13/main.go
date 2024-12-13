package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	dx, dy int
}

type Prize struct {
	x, y int
}

type Machine struct {
	a, b  Button
	prize Prize
}

type Data = []Machine

func parseButton(line string) Button {
	re := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		panic("cannot parse button")
	}

	dx, _ := strconv.Atoi(matches[1])
	dy, _ := strconv.Atoi(matches[2])
	return Button{dx, dy}
}

func parsePrize(line string) Prize {
	re := regexp.MustCompile(`X=(\d+), Y=(\d+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		panic("cannot parse parse")
	}

	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return Prize{x, y}
}

func parseMachine(scanner *bufio.Scanner) Machine {
	buttonA := parseButton(scanner.Text())
	scanner.Scan()
	buttonB := parseButton(scanner.Text())
	scanner.Scan()
	prize := parsePrize(scanner.Text())
	scanner.Scan()
	return Machine{
		a:     buttonA,
		b:     buttonB,
		prize: prize,
	}
}

func parse(filename string) Data {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := Data{}
	for scanner.Scan() {
		machines = append(machines, parseMachine(scanner))
	}

	if err = scanner.Err(); err != nil {
		panic("scanner error")
	}

	return machines
}

func (mach Machine) minTokens() (int, bool) {
	a := mach.a
	b := mach.b
	p := mach.prize

	den := a.dx*b.dy - b.dx*a.dy
	alphaNum := p.x*b.dy - p.y*b.dx
	betaNum := a.dx*p.y - a.dy*p.x

	if alphaNum%den != 0 || betaNum%den != 0 {
		return 0, false
	}
	alpha := alphaNum / den
	beta := betaNum / den

	return 3*alpha + 1*beta, true
}

func p1(data Data) int {
	tokens := 0
	for _, mach := range data {
		tok, ok := mach.minTokens()
		if ok {
			tokens += tok
		}
	}
	return tokens
}

func p2(data Data) int {
	offset := 10000000000000

	tokens := 0
	for _, mach := range data {
		mach.prize.x += offset
		mach.prize.y += offset

		tok, ok := mach.minTokens()
		if ok {
			tokens += tok
		}
	}
	return tokens
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
