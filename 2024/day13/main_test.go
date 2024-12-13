package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	data := parse("test.txt")
	expected := Data{
		Machine{
			a:     Button{94, 34},
			b:     Button{22, 67},
			prize: Prize{8400, 5400},
		},
		Machine{
			a:     Button{26, 66},
			b:     Button{67, 21},
			prize: Prize{12748, 12176},
		},
		Machine{
			a:     Button{17, 86},
			b:     Button{84, 37},
			prize: Prize{7870, 6450},
		},
		Machine{
			a:     Button{69, 23},
			b:     Button{27, 71},
			prize: Prize{18641, 10279},
		},
	}
	if !reflect.DeepEqual(expected, data) {
		t.Fatalf("wrong output, expected: %v, got: %v", expected, data)
	}
}

func TestP1(t *testing.T) {
	data := parse("test.txt")
	output := p1(data)
	expected := 480
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}

func TestP2(t *testing.T) {
	data := parse("test.txt")
	output := p2(data)
	expected := 875318608908
	if output != expected {
		t.Errorf("wrong output, expected: %d, got %d", expected, output)
	}
}
