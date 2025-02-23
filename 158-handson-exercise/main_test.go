package main

import "testing"

func TestFoo(t *testing.T) {
	xs := foo()
	if xs != "foo" {
		t.Errorf("expected %v got %v", "foo", xs)
	}
}
