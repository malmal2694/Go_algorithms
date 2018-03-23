package heapSort

func HeapSort(heap []int) []int {
	var heapSorted []int

	heapSize := len(heap) //Save size of heap
	for i := 0; i < heapSize; i++ {
		heapSorted = append(heapSorted, extractMin(&heap))
	}
	return heapSorted
} // End heapSort func
func extractMin(heap *[]int) int {
	var smallest int = (*heap)[0]
	var j = 0

	for i := 1; i < len(*heap); i++ {
		if (*heap)[i] < smallest {
			smallest = (*heap)[i]
			j = i
		}
	}
	*heap = removeIndex(*heap, j)
	return smallest
}
func removeIndex(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
