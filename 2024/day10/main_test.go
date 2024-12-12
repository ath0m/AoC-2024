package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := parse("test.txt")
	expected := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	if !reflect.DeepEqual(expected, *data) {
		t.Fatalf("wrong output, expected: %v, got: %v", expected, *data)
	}
}

func TestP1(t *testing.T) {
	data := parse("test.txt")
	output := p1(data)
	expected := 36
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}

func TestP2(t *testing.T) {
	data := parse("test.txt")
	output := p2(data)
	expected := 81
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}
