package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data = []int

func parse(filename string) *Data {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("cannot open file")
	}
	line := strings.TrimSpace(string(content))
	nums := []int{}
	for _, word := range strings.Split(line, " ") {
		num, err := strconv.Atoi(word)
		if err != nil {
			panic("cannot convert word to int")
		}
		nums = append(nums, num)
	}
	return &nums
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	if num < 0 {
		panic("argument needs to be non-negative")
	}
	cnt := 0
	for num > 0 {
		num /= 10
		cnt += 1
	}
	return cnt
}

func splitNumber(num int) (int, int) {
	n := countDigits(num)
	if n%2 == 1 {
		panic("num needs to be even")
	}

	base := 1
	for i := 0; i < n/2; i++ {
		base *= 10
	}
	return num / base, num % base
}

func simulate(data *Data, steps int) int {
	line := *data

	for i := 0; i < steps; i++ {
		newLine := []int{}

		for _, val := range line {
			switch {
			case val == 0:
				newLine = append(newLine, 1)
			case (countDigits(val) % 2) == 0:
				first, second := splitNumber(val)
				newLine = append(newLine, first)
				newLine = append(newLine, second)
			default:
				newLine = append(newLine, val*2024)
			}
		}

		line = newLine
	}

	return len(line)
}

type Pair struct {
	fst, snd int
}

func simulateImproved() func(int, int) int {
	cache := map[Pair]int{}
	var sim func(int, int) int
	sim = func(stone int, steps int) int {
		key := Pair{stone, steps}
		val, ok := cache[key]
		if ok {
			return val
		}

		switch {
		case steps == 0:
			cache[key] = 1
		case stone == 0:
			cache[key] = sim(1, steps-1)
		case (countDigits(stone) % 2) == 0:
			first, second := splitNumber(stone)
			cache[key] = sim(first, steps-1) + sim(second, steps-1)
		default:
			cache[key] = sim(stone*2024, steps-1)
		}
		return cache[key]
	}
	return sim
}

func p1(data *Data) int {
	sim := simulateImproved()
	result := 0
	for _, stone := range *data {
		result += sim(stone, 25)
	}
	return result
}

func p2(data *Data) int {
	sim := simulateImproved()
	result := 0
	for _, stone := range *data {
		result += sim(stone, 75)
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
