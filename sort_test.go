package sort

import "testing"

type MockSort struct {
	sortableValue int64
}

func (s *MockSort) SortableValue() int64 {
	return s.sortableValue
}

func expectArraySorted(t *testing.T, expected []Sort, response []Sort) {
	for i := 0; i < len(expected); i++ {
		if response[i].SortableValue() != expected[i].SortableValue() {
			t.Errorf("Expect %d, got %d", expected[i], response[i])
		}
	}
}
