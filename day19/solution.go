package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

//go:embed sample.txt
var sample string

func parseInput(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")

	patterns := strings.Split(lines[0], ", ")
	designs := []string{}

	for _, line := range lines[2:] {
		if len(line) > 0 {
			designs = append(designs, line)
		}
	}

	return patterns, designs
}

type TrieNode struct {
	children [26](*TrieNode)
	isWord   bool
	id       int
}

func (node *TrieNode) add(word string, id int) {
	for _, c := range word {
		i := int(c) - int('a')
		if node.children[i] == nil {
			node.children[i] = &TrieNode{}
		}
		node = node.children[i]
	}
	node.isWord = true
	node.id = id
}

func buildTrie(patterns []string) *TrieNode {
	root := TrieNode{}
	for i, pattern := range patterns {
		root.add(pattern, i)
	}
	return &root
}

func findWithCount(word string, root *TrieNode, cache map[string]int) int {
	if len(word) == 0 {
		return 1
	}

	val, ok := cache[word]
	if ok {
		return val
	}

	cache[word] = 0
	node := root
	for i, c := range word {
		index := int(c) - int('a')
		if node.children[index] == nil {
			break
		}
		node = node.children[index]
		if node.isWord {
			cache[word] += findWithCount(word[i+1:], root, cache)
		}
	}
	return cache[word]
}

func findWithCountSlow(word string, patterns []string, cache map[string]int) int {
	if len(word) == 0 {
		return 1
	}

	val, ok := cache[word]
	if ok {
		return val
	}

	cache[word] = 0
	for _, pattern := range patterns {
		n := len(pattern)
		if n > len(word) {
			break
		}
		if word[:n] == pattern {
			cache[word] += findWithCountSlow(word[n:], patterns, cache)
		}
	}
	return cache[word]
}

func partOne(patterns, designs []string) int {
	root := buildTrie(patterns)
	cache := map[string]int{}

	cnt := 0
	for _, design := range designs {
		if findWithCount(design, root, cache) > 0 {
			cnt += 1
		}
	}

	return cnt
}

func partTwo(patterns, designs []string) int {
	root := buildTrie(patterns)
	cache := map[string]int{}

	cnt := 0
	for _, design := range designs {
		cnt += findWithCount(design, root, cache)
	}
	return cnt
}
