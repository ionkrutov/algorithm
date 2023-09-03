# algorithm

## Heap
Heap data structure ([Wiki](https://en.wikipedia.org/wiki/Heap_(data_structure))) is realised with golang generics.
Attractiveness of this package is it simplicity to use,
because need to define only one function -- Comparator.

See the [Documentation](https://pkg.go.dev/github.com/ionkrutov/algorithm/heap)

### Examples
A simple heap sort can be realised if we will extract top element one by one and put it in another array.

```bash
go get -u github.com/ionkrutov/algorithm/heap
```

```go
package main

import (
	"fmt"

	"github.com/ionkrutov/algorithm/heap"
)

func main() {
	h := heap.Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8}...)
	fmt.Println(h.GetHeap())

	// heap sort
	sorted := make([]int, 0, h.Size())
	for h.Size() > 0 {
		sorted = append(sorted, h.ExtractTop())
	}
	fmt.Println(sorted)
}
```
result
```bash
[8 5 7 4 1 6 3 2]
[8 7 6 5 4 3 2 1]
```


We can use a custom structure

```go
package main

import (
	"fmt"

	"github.com/ionkrutov/algorithm/heap"
)

func main() {
	type Person struct {
		Name   string
		Age    int
		Salary float32
	}

	// min heap by salary
	h := heap.Heap[Person]{Comparator: func(p1, p2 Person) bool { return p1.Salary < p2.Salary }}
	h.Insert(Person{"Anna", 32, 122.34})
	h.Insert(Person{"Peter", 45, 212.4})
	h.Insert(Person{"Igori", 23, 74.2})
	h.Insert(Person{"Vasilii", 23, 98.3})
	h.Insert(Person{"Hellen", 24, 49.3})

	sorted := make([]Person, 0, h.Size())
	for h.Size() > 0 {
		sorted = append(sorted, h.ExtractTop())
	}
	fmt.Println(sorted)
}
```

result
```bash
[{Hellen 24 49.3} {Igori 23 74.2} {Vasilii 23 98.3} {Anna 32 122.34} {Peter 45 212.4}]
```