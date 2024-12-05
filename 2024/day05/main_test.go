package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	rules, updates := parse("test.txt")

	expected_rules := []rule{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}
	if !reflect.DeepEqual(expected_rules, rules) {
		t.Errorf("Wrong rules, want: %v, got: %v", expected_rules, rules)
	}

	expected_updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	if !reflect.DeepEqual(expected_updates, updates) {
		t.Errorf("Wrong updates, want: %v, got: %v", expected_updates, updates)
	}
}

func TestP1(t *testing.T) {
	rules, updates := parse("test.txt")
	result := p1(&rules, &updates)
	expected := 143
	if result != expected {
		t.Errorf("Wrong P1, want: %d, got: %d", expected, result)
	}
}

func TestP2(t *testing.T) {
	rules, updates := parse("test.txt")
	result := p2(&rules, &updates)
	expected := 123
	if result != expected {
		t.Errorf("Wrong P2, want: %d, got: %d", expected, result)
	}
}
