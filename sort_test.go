package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	emptyList := make([]Entry, 50)
	seed := rand.New(rand.NewSource(time.Now().UnixNano())) // ensures randomness
	for i, _ := range emptyList {
		emptyList[i].Priority = seed.Intn(10000) // this allows for uniquity among the entries
	}
	sorted := Sort(emptyList)
	var error int
	for i := range sorted {
		if sorted[i].Priority == emptyList[i].Priority {
			error++
		}
	}
	if error > 5 { // our MOE can't be greater than 1/10
		t.Error("List not randomly sorted.")
	}
}

func TestRandomPrioritize(t *testing.T) {
	emptyList := make([]Entry, 5)
	seed := rand.New(rand.NewSource(time.Now().UnixNano())) // ensures randomness
	for i, _ := range emptyList {
		emptyList[i].Priority = seed.Intn(500)
	}
	sorted := Sort(emptyList)
	prioritized := Prioritize(sorted)
	lastNumber := 501 // using 501 because this will always be largest number
	for _, v := range prioritized {
		if lastNumber < v.Priority {
			t.Error("Priority numbers not properly sorted")
		}
		lastNumber = v.Priority
	}
}

func TestAdmit(t *testing.T) {
	emptyList := make([]Entry, 5)
	sorted := Sort(emptyList)
	admitted := Admit(sorted, 1)
	if admitted[0].Status != ADMITTED && admitted[0].Status != WAITLISTED {
		t.Error("Improper amount of entries admitted.")
	}
}
