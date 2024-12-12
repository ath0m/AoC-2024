package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := parse("test.txt")
	expected := Data{125, 17}
	if !reflect.DeepEqual(expected, *data) {
		t.Fatalf("wrong output, expected: %v, got: %v", expected, *data)
	}
}

func TestP1(t *testing.T) {
	data := parse("test.txt")
	output := p1(data)
	expected := 55312
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}

func TestP2(t *testing.T) {
	data := parse("test.txt")
	output := p2(data)
	expected := 0
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}
