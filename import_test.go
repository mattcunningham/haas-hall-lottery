package main

import "testing"

func TestCSVtoEntries(t *testing.T) {
	in := `Grade, Sibling, Faculty, First name
8,,,Matthew
9,YES,,Davis
10,,YES,Jade`
	entries := CSVtoEntries(in)
	expectedOutput := []Entry{
		Entry{Status: 0, Grade: 8, Priority: 0, Info: map[string]string{"First name": "Matthew"}},
		Entry{Status: 0, Grade: 9, Priority: 1, Info: map[string]string{"First name": "Davis"}},
		Entry{Status: 0, Grade: 10, Priority: 2, Info: map[string]string{"First name": "Jade"}},
	}
	for i, entry := range entries {
		expectedOutputEntry := expectedOutput[i]
		if entry.Status != expectedOutputEntry.Status {
			t.Error("Expected status does not match actual status.")
		}
		if entry.Grade != expectedOutputEntry.Grade {
			t.Error("Expected grade does not match actual grade.")
		}
		if entry.Priority != expectedOutputEntry.Priority {
			t.Error("Expected priority does not match actual priority.")
		}

	}
}
