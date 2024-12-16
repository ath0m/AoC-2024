package main

import (
	"slices"
	"testing"
)

func TestParse(t *testing.T) {
	diskmap := parse("test.txt")
	expected := []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2}
	if !slices.Equal(*diskmap, expected) {
		t.Errorf("wrong output, expected: %v, got %v", expected, *diskmap)
	}
}

func TestExpandDiskmap(t *testing.T) {
	diskmap := []int{1, 2, 3, 4, 5}
	blocks := expandDiskmap(&diskmap)
	expected := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	if !slices.Equal(*blocks, expected) {
		t.Errorf("wrong output, expected: %v, got: %v", expected, *blocks)
	}
}

func TestP1(t *testing.T) {
	diskmap := parse("test.txt")
	output := p1(diskmap)
	expected := 1928
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}

func TestP2(t *testing.T) {
	diskmap := parse("test.txt")
	output := p2(diskmap)
	expected := 2858
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}
