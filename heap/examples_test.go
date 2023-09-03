package heap

import "fmt"

func ExampleHeap_Insert() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert(5)
	h.Insert(7)
	h.Insert([]int{10, 12, 14, 16}...)

	fmt.Println(h.GetHeap())

	// Output:
	// [16 14 10 12 5 7]
}

func ExampleHeap_Top() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert(5)
	h.Insert(7)
	h.Insert([]int{10, 12, 14, 16}...)
	fmt.Println(h.Top())
	fmt.Println(h.Size())

	// Output:
	// 16
	// 6
}

func ExampleHeap_ExtractTop() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert(5)
	h.Insert(7)
	h.Insert([]int{10, 12, 14, 16}...)
	fmt.Println(h.ExtractTop())
	fmt.Println(h.Size())
	// Output:
	// 16
	// 5
}

func ExampleHeap_Remove() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert(5)
	h.Insert(7)
	h.Insert([]int{10, 12, 14, 16}...)
	h.Remove(2)
	fmt.Println(h.GetHeap())

	// Output:
	// [16 14 7 12 5]
}

func ExampleHeap_ChangePriority() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert([]int{5, 7, 10, 12, 14, 16}...)
	fmt.Println(h.GetHeap())
	h.ChangePriority(0, 2)
	fmt.Println(h.GetHeap())

	// Output:
	// [16 14 10 12 7 5]
	// [14 12 10 2 7 5]
}

func ExampleHeap_GetHeap() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert([]int{5, 7, 10, 12, 14, 16}...)
	fmt.Println(h.GetHeap())

	// Output:
	// [16 14 10 12 7 5]
}

func ExampleHeap_Size() {
	// create max heap
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert([]int{5, 7, 10, 12, 14, 16}...)
	fmt.Println(h.Size())

	// Output:
	// 6
}

func ExampleCustomStruct() {
	type Person struct {
		Name   string
		Age    int
		Salary float32
	}

	// min heap by salary
	h := Heap[Person]{Comparator: func(p1, p2 Person) bool { return p1.Salary < p2.Salary }}
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

	// Output:
	// [{Hellen 24 49.3} {Igori 23 74.2} {Vasilii 23 98.3} {Anna 32 122.34} {Peter 45 212.4}]
}
