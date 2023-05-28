package sorting

import (
	"math/rand"
	"time"
)

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func MergeSorting(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSorting(arr[:mid])
	right := MergeSorting(arr[mid:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(arr))
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]

	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	QuickSort(arr[:i])
	QuickSort(arr[i+1:])

	return arr
}

func SelectionSort(arr []int) []int {
	sortedArr := make([]int, len(arr))
	i, j := 0, 0
	for len(arr) != 0 {
		sortedArr[i] = arr[0]
		j = 0
		for k := 0; k < len(arr); k++ {
			if sortedArr[i] > arr[k] {
				sortedArr[i] = arr[k]
				j = k
			}
		}
		i++
		arr = append(arr[j+1:], arr[:j]...)
	}
	return sortedArr
}
