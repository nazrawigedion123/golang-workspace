package main

import "testing"

func TestDoMath(t *testing.T) {
	total := doMath(5, 5, add)
	if total != 10 {
		t.Errorf("expected %d got %d", 10, total)
	}
	total1 := doMath(5, 5, substract)
	if total1 != 0 {
		t.Errorf("expected %d got %d", 0, total)
	}
}

func TestSubstract(t *testing.T) {
	total := substract(5, 5)
	if total != 0 {
		t.Errorf("expected %d got %d", 0, total)
	}
}

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("expected %d got %d", 10, total)
	}
}
