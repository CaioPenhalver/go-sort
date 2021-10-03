package sort

func InsertionSort(array []Sort) []Sort {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j > -1 && array[j].SortableValue() > key.SortableValue() {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
	return array
}
