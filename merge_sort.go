package sort

import "fmt"

func MergeSort(array []Sort) []Sort {
	if len(array) <= 1 {
		return array
	}
	mergeSort(array, 0, len(array)-1)
	return array
}

func mergeSort(array []Sort, left int, right int) {
	if left < right {
		middle := (left + right) / 2

		mergeSort(array, left, middle)
		mergeSort(array, middle+1, right)

		merge(array, left, middle, right)
	}
}

func merge(array []Sort, left int, middle int, right int) {
	leftTempSize := middle - left + 1
	rightTempSize := right - middle

	leftArray := make([]Sort, leftTempSize)
	rightArray := make([]Sort, rightTempSize)

  // TODO
  // leftArray = array[left..leftTempSize]
	for i := 0; i < leftTempSize; i++ {
		leftArray[i] = array[left+i]
	}

	for j := 0; j < rightTempSize; j++ {
		rightArray[j] = array[middle+j+1]
	}

	i, j := 0, 0
  k := left
	for i < leftTempSize && j < rightTempSize {
		if leftArray[i].SortableValue() <= rightArray[j].SortableValue() {
			array[k] = leftArray[i]
			i++
		} else {
			array[k] = rightArray[j]
			j++
		}
		k++
	}

	for i < leftTempSize {
		array[k] = leftArray[i]
		i++
		k++
	}

	for j < rightTempSize {
		array[k] = rightArray[j]
		j++
		k++
	}
}

func logArray(a []Sort) {
	for _, v := range a {
		fmt.Printf("%d ", v.SortableValue())
	}
	fmt.Printf("\n")
}
