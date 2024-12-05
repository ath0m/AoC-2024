package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func parse(filename string) ([]rule, [][]int) {
	rules := []rule{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatal("The input string is not in the expected format.")
		}

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting first number: %v\n", err)
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting second number: %v\n", err)
		}

		rules = append(rules, rule{before, after})
	}

	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		update := make([]int, len(nums))

		for i, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error converting number: %v\n", err)
			}
			update[i] = val
		}

		updates = append(updates, update)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, updates
}

func verify(rules *[]rule, update *[]int) bool {
	indices := map[int]int{}
	for i, num := range *update {
		indices[num] = i
	}

	for _, rle := range *rules {
		l, ok := indices[rle.before]
		if !ok {
			continue
		}
		r, ok := indices[rle.after]
		if !ok {
			continue
		}
		if l > r {
			return false
		}
	}
	return true
}

func p1(rules *[]rule, updates *[][]int) int {
	result := 0
	for _, update := range *updates {
		if verify(rules, &update) {
			result += update[len(update)/2]
		}
	}
	return result
}

func fix(rules *[]rule, update *[]int) []int {
	cnt := map[int]int{}
	graph := map[int]*[]int{}
	for _, val := range *update {
		graph[val] = &[]int{}
		cnt[val] = 0
	}

	for _, rle := range *rules {
		fst, ok := graph[rle.before]
		if !ok {
			continue
		}
		_, ok = graph[rle.after]
		if !ok {
			continue
		}
		*fst = append(*fst, rle.after)
		cnt[rle.after] += 1
	}

	q := []int{}
	for k, v := range cnt {
		if v == 0 {
			q = append(q, k)
		}
	}

	fixed := []int{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		fixed = append(fixed, v)
		deps, _ := graph[v]
		for _, u := range *deps {
			cnt[u] -= 1
			if cnt[u] == 0 {
				q = append(q, u)
			}
		}
	}

	return fixed
}

func p2(rules *[]rule, updates *[][]int) int {
	result := 0
	for _, update := range *updates {
		if !verify(rules, &update) {
			fixed := fix(rules, &update)
			result += fixed[len(fixed)/2]
		}
	}
	return result
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Argument needs to be passed")
	}
	rules, updates := parse("input.txt")
	switch args[1] {
	case "p1":
		fmt.Println("p1:", p1(&rules, &updates))
	case "p2":
		fmt.Println("p2:", p2(&rules, &updates))
	}
}
