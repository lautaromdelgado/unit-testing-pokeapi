package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	result := Add(20, 2)
	expected := 22
	if result != expected {
		t.Errorf("se esperaba %d, pero se obtuvo %d", expected, result)
	}
}
