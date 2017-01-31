package main

import (
	"math/rand"
	"time"
)

const (
	NIL = iota
	WAITLISTED
	ADMITTED
)

type Entry struct {
	Status   int               // admitted or waitlisted
	Priority int               // priority status
	Grade    int               // grade of student
	Info     map[string]string // unnecessary to formally store private info
}

// this is first step
// randomly sorts entries, no acceptances/waitlists are determined here
func Sort(allEntries []Entry) []Entry {
	seed := rand.New(rand.NewSource(time.Now().UnixNano())) // ensures randomness
	perm := seed.Perm(len(allEntries))                      // permutation of all numbers [0, len(allEntries))
	sorted := make([]Entry, len(allEntries))                // creating new array that will be sorted
	for i, v := range perm {
		sorted[i] = allEntries[v] // for every index, a random entry is selected
	}
	return sorted
}

// The entries are sorted so higher priority numbers
// are placed at the front of the list. Typically, there will
// only be two priority numbers for faculty and siblings.
// Additional numbers depend on school needs.
func Prioritize(allEntries []Entry) []Entry {
	var priority, fullList []Entry
	for _, v := range allEntries {
		if v.Priority > 0 {
			priority = append(priority, v) // this priority list will precede the regular list
		} else {
			fullList = append(fullList, v) // if not priority, go to regular list
		}
	}
	priority = MergeSort(priority)
	return append(priority, fullList...)
}

// part of merge sort algorithm; merges left/right halves of slice
func Merge(l, r []Entry) []Entry {
	ret := make([]Entry, 0, len(l)+len(r)) // return value
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0].Priority >= r[0].Priority {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

// traditional mergesort sorting algorithm
// due to its stability, it's the best choice
func MergeSort(entries []Entry) []Entry {
	if len(entries) <= 1 {
		return entries
	}
	n := len(entries) / 2
	l := MergeSort(entries[:n])
	r := MergeSort(entries[n:])
	return Merge(l, r)
}

// given a cap of int size, will add to the []Entry struct if scholar is admitted or not
func Admit(allEntries []Entry, limit int) []Entry {
	for i := range allEntries {
		if i < limit {
			allEntries[i].Status = ADMITTED
		} else {
			allEntries[i].Status = WAITLISTED
		}
	}
	return allEntries
}
