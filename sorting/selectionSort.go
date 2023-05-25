package sorting

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
