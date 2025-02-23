package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(6, 5)
	if total != 10 {
		t.Errorf("sum was incorrect got %d but expected %d", total, 10)
	}
}
