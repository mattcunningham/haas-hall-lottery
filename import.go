package main

import (
	"encoding/csv"
	"io"
	"strings"
)

func CSVtoEntries(data string) []Entry {
	r := csv.NewReader(strings.NewReader(data))
	var recordNames []string
	for {
		record, err := r.ReadAll()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // this needs to be a better error, prob return to tell user
		}
		if len(recordNames) == 0 {
			recordNames = record
		} else {

		}
	}
}
