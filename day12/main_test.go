package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := parse("test.txt")
	expected := Data{
		"RRRRIICCFF",
		"RRRRIICCCF",
		"VVRRRCCFFF",
		"VVRCCCJFFF",
		"VVVVCJJCFE",
		"VVIVCCJJEE",
		"VVIIICJJEE",
		"MIIIIIJJEE",
		"MIIISIJEEE",
		"MMMISSJEEE",
	}
	if !reflect.DeepEqual(expected, *data) {
		t.Fatalf("wrong output, expected: %v, got: %v", expected, *data)
	}
}

func TestP1(t *testing.T) {
	data := parse("test.txt")
	output := p1(data)
	expected := 1930
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}

func TestP2(t *testing.T) {
	data := parse("test.txt")
	output := p2(data)
	expected := 1206
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}
