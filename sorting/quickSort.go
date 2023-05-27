package sorting

import (
	"math/rand"
	"time"
)

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
