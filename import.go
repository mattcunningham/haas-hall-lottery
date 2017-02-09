// written by Matthew Cunningham
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
	var priority, grade, lotteryID int
	var entries []Entry
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err) // this needs to be a better error, prob return to tell user
		}
		if len(recordNames) == 0 {
			recordNames, priority, grade, lotteryID = SetNames(record)
		} else {
			entries = append(entries, CSVtoEntryMap(recordNames, record, priority, grade, lotteryID))
		}
	}
	return entries
}

// this gets the index of the main fields
func SetNames(names []string) (recordNames map[int]string, priority int, grade int, lotteryID int) {
	recordNames = make(map[int]string, len(names))
	for i, name := range names {
		name = strings.TrimSpace(name)
		recordNames[i] = name
		if name == "Priority" {
			priority = i
		} else if name == "Grade" {
			grade = i
		} else if name == "Lottery ID" {
			lotteryID = i
		}
	}
	return recordNames, priority, grade, lotteryID
}

func CSVtoEntryMap(recordNames map[int]string, fullRecord []string, priority int, grade int, lotteryID int) Entry {
	var entry Entry
	entry.Info = make(map[string]string, len(fullRecord)-3)
	for i, record := range fullRecord {
		if i != priority && i != grade && i != lotteryID {
			entry.Info[recordNames[i]] = record
		} else if i == priority && strings.ToLower(strings.Trim(record, " ")) == "faculty" { 
			entry.Priority = 2 // higher priority
		} else if i == priority && strings.ToLower(strings.Trim(record, " ")) == "sibling" {
			entry.Priority = 1 // high priority
		} else if i == grade {
			entry.Grade, _ = strconv.Atoi(record) // will error if value isn't numerical. TODO: deal with this
		} else if i == lotteryID {
			entry.LotteryID = record
		}
	}
	return entry
}
