package main

import (
	"rand"
)

const (
	NIL = iota
	WAITLISTED
	ADMITTED
)

type Entry struct {
	Status  int               // admitted or waitlisted
	Faculty bool              // gets first priority
	Sibling bool              // second priority
	Info    map[string]string // unnecessary to formally store private info
}

// randomly sorts entries, no acceptances/waitlists are determined here
func Sort(allEntries []Entry) []Entry {
	seed := rand.New()                       // ensuring random seed
	perm := seed.Perm(len(allEntries))       // permutation of all numbers [0, len(allEntries))
	sorted := make([]Entry, len(allEntries)) // creating new array that will be sorted
	for i, v := range perm {
		sorted[i] = allEntries[v] // for every index, a random entry is selected
	}
	return sorted
}

// provide faculty member children with ultimate priority
// provide siblings with second priority
// no more priority after that
func Prioritize(allEntries []Entry) []Entry {
	var (
		facultyPriority []Entry
		siblingPriority []Entry
	)
	for i, v := range allEntries {
		if v.Faculty {
			facultyPriority = append(facultyPriority, v)
			allEntries = append(allEntries[:i], allEntries[i+1:])
		} else if v.Sibling {
			siblingPriority = append(facultyPriority, v)
			allEntries = append(allEntries[:i], allEntries[i+1:])
		}
	}
	allEntries = append(facultyPriority, append(siblingPriority, allEntries...)...) // merging slices together is tricky. the dots make it a variadic function. probably a better way to do this
	return allEntries
}

// given a cap of int size, will add to the []Entry struct if scholar is admitted or not
func Admit(allEntries []Entry, limit int) []Entry {
	for i, v := range allEntries {
		if i < limit {
			v.Status = ADMITTED
		} else {
			v.Status = WAITLISTED
		}
	}
}
