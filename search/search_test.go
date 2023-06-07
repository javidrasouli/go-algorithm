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
