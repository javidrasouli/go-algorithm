package search

func LinearSearch(arr []int, number int) (foundNumber int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			foundNumber = arr[i]
			break
		}
	}
	return
}
