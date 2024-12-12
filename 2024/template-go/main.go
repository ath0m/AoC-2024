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
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic("scanner error")
	}

	return &lines
}

func p1(data *Data) int {
	return 0
}

func p2(data *Data) int {
	return 0
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
