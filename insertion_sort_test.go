package sort

import "testing"

func TestInsertionSort(t *testing.T) {
  objects := []Sort{
    &MockSort{sortableValue: 298},
    &MockSort{sortableValue: 4},
    &MockSort{sortableValue: 97},
    &MockSort{sortableValue: 1},
    &MockSort{sortableValue: 199},
    &MockSort{sortableValue: 78},
    &MockSort{sortableValue: 100},
    &MockSort{sortableValue: 8},
    &MockSort{sortableValue: 7},
  }
  expectedObjects := []Sort{
    &MockSort{sortableValue: 1},
    &MockSort{sortableValue: 4},
    &MockSort{sortableValue: 7},
    &MockSort{sortableValue: 8},
    &MockSort{sortableValue: 78},
    &MockSort{sortableValue: 97},
    &MockSort{sortableValue: 100},
    &MockSort{sortableValue: 199},
    &MockSort{sortableValue: 298},
  }

	InsertionSort(objects)

	expectArraySorted(t, expectedObjects, objects)
}
