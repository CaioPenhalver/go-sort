# go-sort

## Install

```bash
go get github.com/CaioPenhalver/go-sort
```

## Features

- bubble sort
- insertion sort
- merge sort


## Example

```go
import(
  mySort "github.com/CaioPenhalver/go-sort"
) 

type Object struct {
    sortableValue int64
}

func (o *Object) SortableValue() int64 {
    return o.sortableValue
}
...
objs := []mySort.Sort {
    &Object{sortableValue: 34},
    &Object{sortableValue: 1},
    &Object{sortableValue: 89},
    &Object{sortableValue: 99},
}

mySort.BubbleSort(objs)
...
```

## Development
wip

## Contributing
wip
