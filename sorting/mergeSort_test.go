package sorting

import (
	"testing"
)

func TestMergeSorting(t *testing.T) {
	mainArr := []int{1, 4, 7, 6, 9, 8, 5, 3, 2}
	sortArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := MergeSorting(mainArr)
	if len(result) != len(sortArr) {
		t.Errorf("array len must be %q, but is %q", len(sortArr), len(result))
	}
	for i := 0; i < len(result); i++ {
		if result[i] != sortArr[i] {
			t.Errorf("number in index %q must be %q, but is %q", i, sortArr[i], result[i])
		}
	}
}
