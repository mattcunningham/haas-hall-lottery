package lottery

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

// The data coming should be a large string in CSV format. For example:
// Grade, Sibling, Faculty,  First name, Last name
// 12, , , Matthew, Cunningham
// 12, Yes, , Jade, DeSpain
func CSVtoEntries(data string) []Entry {
	r := csv.NewReader(strings.NewReader(data))
	var recordNames map[int]string
	var sibling, faculty, grade int
	var entries []Entry
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // this needs to be a better error, prob return to tell user
		}
		if len(recordNames) == 0 {
			recordNames, sibling, faculty, grade = SetNames(record)
		} else {
			entries = append(entries, CSVtoEntryMap(recordNames, record, sibling, faculty, grade))
		}
	}
	return entries
}

func SetNames(names []string) (recordNames map[int]string, sibling int, faculty int, grade int) {
	recordNames = make(map[int]string, len(names))
	for i, name := range names {
		name = strings.TrimSpace(name)
		recordNames[i] = name
		if name == "Sibling" {
			sibling = i
		} else if name == "Faculty" {
			faculty = i
		} else if name == "Grade" {
			grade = i
		}
	}
	return recordNames, sibling, faculty, grade
}

func CSVtoEntryMap(recordNames map[int]string, fullRecord []string, sibling int, faculty int, grade int) Entry {
	var entry Entry
	entry.Info = make(map[string]string, len(fullRecord)-3)
	for i, record := range fullRecord {
		if i != sibling && i != faculty && i != grade {
			entry.Info[recordNames[i]] = record
		} else if i == faculty && strings.ToLower(strings.Trim(record, " ")) == "yes" { // If this is the faculty record and it has ANY value, priority given next line
			entry.Priority = 2 // higher priority
		} else if i == sibling && strings.ToLower(strings.Trim(record, " ")) == "yes" { // If this is the sibling record and it has ANY value, priority given next line
			entry.Priority = 1 // high priority
		} else if i == grade {
			entry.Grade, _ = strconv.Atoi(record) // will error if value isn't numerical. TODO: deal with this
		}
	}
	return entry
}
