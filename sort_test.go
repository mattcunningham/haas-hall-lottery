package main

import "testing"

func TestSort(t *testing.T) {
	emptyList := make([]Entry, 100)
	sorted := Sort(emptyList)
	if sorted == emptyList {
		t.Error("List was not randomly sorted.")
	}
}
