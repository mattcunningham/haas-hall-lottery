package main

import "testing"

func TestCSVtoEntries(t *testing.T) {
	in := `Grade, Sibling, Faculty, Fayetteville
8, No, No, Yes
9, Yes, No, No`
	entries := CSVToEntries(in)
	expectedOutput := []Entry{
		Entry{Status: 0, Priority: 0, Info: map[string]string{"Fayetteville": "Yes"}},
		Entry{Status: 0, Priority: 1, Info: map[string]string{"Fayetteville": "No"}},
	}
	if entries != expectedOutput {
		t.Error("CSVToEntries did not return the expected output.")
	}
)
