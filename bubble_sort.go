package sort

func BubbleSort(array []Sort) {
	length := len(array) - 1

	for i := 0; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if array[j].SortableValue() > array[j+1].SortableValue() {
				obj := array[j]
				array[j] = array[j+1]
				array[j+1] = obj
			}
		}
	}
}
