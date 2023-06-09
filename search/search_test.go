package search

import "testing"

func TestLinearSearch(t *testing.T) {
	mainArr := []int{1, 4, 7, 6, 9, 8, 5, 3, 2}
	numberToSearch := 8
	foundNumber := LinearSearch(mainArr, numberToSearch)
	if numberToSearch != foundNumber {
		t.Errorf("found number must be %q, but is %q", numberToSearch, foundNumber)
	}
}

func TestBinarySearch(t *testing.T) {
	mainArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numberToSearch := 8
	foundNumber := BinarySearch(mainArr, numberToSearch)
	if numberToSearch != foundNumber {
		t.Errorf("found number must be %q, but is %q", numberToSearch, foundNumber)
	}
}

func TestDepthFirstSearch(t *testing.T) {
	graph := map[string][]string{
		"A": []string{"B", "G"},
		"B": []string{"C", "D", "E"},
		"C": []string{},
		"D": []string{},
		"E": []string{"F"},
		"F": []string{},
		"G": []string{"H"},
		"H": []string{"I"},
		"I": []string{},
	}
	Word := "F"
	found := DepthFirstSearch(graph, "A", Word)
	if found != Word {
		t.Errorf("found word must be %q, but is %q", Word, found)
	}
}
