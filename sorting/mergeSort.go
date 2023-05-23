package sorting

func MergeSorting(arr []int, start int, end int) []int {
	if end > start {
		i := (start + end) / 2
		arr = MergeSorting(arr, start, i)
		arr = MergeSorting(arr, i+1, end)
		arr = Merge(arr[i:], arr[:i], i)
	}
	return arr
}
func Merge(arr1, arr2 []int, i int) []int {
	var sorted []int
	m := 0
	k := 0
	arrLen := len(arr1) + len(arr2)
	for i := 0; i < arrLen; i++ {
		if m >= len(arr1) {
			sorted = append(sorted, arr2[k:]...)
			return sorted
		} else if k >= len(arr2) {
			sorted = append(sorted, arr1[m:]...)
			return sorted
		}
		if arr1[m] < arr2[k] {
			sorted = append(sorted, arr1[m])
			m = m + 1
			continue
		}
		if arr1[m] > arr2[k] {
			sorted = append(sorted, arr2[k])
			k = k + 1
			continue
		}
	}
	return sorted
}
