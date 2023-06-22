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

func BinarySearch(arr []int, number int) (foundNumber int) {
	index := len(arr) / 2
	rightArr := arr[:index]
	leftArr := arr[index:]
	if len(rightArr) == 1 && rightArr[0] == number {
		foundNumber = rightArr[0]
		return
	}
	if len(leftArr) == 1 && leftArr[0] == number || leftArr[0] == number {
		foundNumber = leftArr[0]
		return
	}
	if rightArr[len(rightArr)-1] == number {
		foundNumber = rightArr[len(rightArr)-1]
	}
	if rightArr[len(rightArr)-1] > number {
		foundNumber = BinarySearch(rightArr, number)
	}
	if leftArr[0] < number {
		foundNumber = BinarySearch(leftArr, number)
	}
	return
}

func DepthFirstSearch(graph map[string][]string, startNode string, word string) (found string) {
	if startNode == word {
		found = startNode
		return
	}
	for _, s := range graph[startNode] {
		if s == word {
			found = s
			return
		}
		if s != startNode {
			found = DepthFirstSearch(graph, s, word)
			if found == word {
				return
			}
		}
	}
	return
}
