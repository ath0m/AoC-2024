package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	grid := parse("test.txt")
	expected := [][]byte{
		[]byte("....#....."),
		[]byte(".........#"),
		[]byte(".........."),
		[]byte("..#......."),
		[]byte(".......#.."),
		[]byte(".........."),
		[]byte(".#..^....."),
		[]byte("........#."),
		[]byte("#........."),
		[]byte("......#..."),
	}
	if !reflect.DeepEqual(*grid, expected) {
		t.Errorf("parse output is invalide, want=%v, got=%v", expected, grid)
	}
}

func TestFindGuard(t *testing.T) {
	grid := parse("test.txt")
	i, j, di, dj := findGuard(grid)
	if i != 6 || j != 4 || di != -1 || dj != 0 {
		t.Errorf("found wrong guard position, want: (%d, %d, %d, %d), got: (%d, %d, %d, %d)", 6, 4, -1, 0, i, j, di, dj)
	}
}

func TestP1(t *testing.T) {
	grid := parse("test.txt")
	output := p1(grid)
	expected := 41
	if output != expected {
		t.Errorf("p1 output is invalid, want=%d, got=%d", expected, output)
	}
}

func TestP2(t *testing.T) {
	grid := parse("test.txt")
	output := p2(grid)
	expected := 6
	if output != expected {
		t.Errorf("p2 output is invalid, want=%d, got=%d", expected, output)
	}
}
