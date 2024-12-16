package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	expected := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	output := parse("test1.txt")
	if !reflect.DeepEqual(expected, *output) {
		t.Fatalf("wrong output from parse, expected: %v, got: %v", expected, *output)
	}
}

func TestP1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"test2.txt", 2},
		{"test3.txt", 4},
		{"test4.txt", 4},
		{"test5.txt", 14},
	}

	for i, tt := range tests {
		grid := parse(tt.filename)
		output := p1(grid)
		if output != tt.expected {
			t.Fatalf("[%d] wrong output from p1, expected: %d, got: %d", i, tt.expected, output)
		}
	}
}

func TestExtractAntennas(t *testing.T) {
	expected := map[rune][]Point{
		'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
		'A': {{5, 6}, {8, 8}, {9, 9}},
	}
	grid := parse("test1.txt")
	antennas := extractAntennas(grid)
	if !reflect.DeepEqual(expected, *antennas) {
		t.Fatalf("invalid extracted antennas, expected: %d, got: %d", expected, *antennas)
	}
}

func TestP2(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"test5.txt", 34},
		{"test6.txt", 9},
	}

	for i, tt := range tests {
		grid := parse(tt.filename)
		output := p2(grid)
		if output != tt.expected {
			t.Fatalf("[%d] wrong output from p2, expected: %d, got: %d", i, tt.expected, output)
		}
	}
}
