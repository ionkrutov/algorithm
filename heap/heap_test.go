package heap

import (
	"fmt"
	"testing"
)

func compareArrays[T comparable](elems ...T) error {
	res := len(elems) % 2
	_ = res
	if len(elems)%2 != 0 {
		return fmt.Errorf("the len of elements aren't equal")
	}
	mid := len(elems) / 2
	for i := 0; i < mid; i++ {
		if elems[i] != elems[i+mid] {
			return fmt.Errorf("%v != %v, for i = %d", elems[i], elems[i+mid], i)
		}
	}
	return nil

}

func TestCompareArrays(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 4, 5, 4}
	combine := append(arr1, arr2...)
	if err := compareArrays(combine...); err == nil {
		t.Error(err)
	}
}

func TestMaxHeap(t *testing.T) {
	h := Heap[int]{}
	h.Comparator = func(i1, i2 int) bool {
		return i1 > i2
	}

	h.Insert([]int{1, 2, 3, 4, 5}...)

	actual := make([]int, 0)
	for h.Size() != 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{5, 4, 3, 2, 1}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestMinHeap(t *testing.T) {
	h := Heap[int]{}
	h.Comparator = func(i1, i2 int) bool {
		return i1 < i2
	}

	h.Insert([]int{5, 4, 3, 2, 1}...)

	actual := make([]int, 0)
	for h.Size() != 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{1, 2, 3, 4, 5}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestHeapMethods(t *testing.T) {
	h := Heap[int]{}
	h.Comparator = func(i1, i2 int) bool {
		return i1 > i2
	}

	h.Insert([]int{5, 4, 3, 2, 1}...)
	actual := make([]int, 0)

	for h.Size() != 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{5, 4, 3, 2, 1}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestGetHeap(t *testing.T) {
	h := Heap[float32]{Comparator: func(f1, f2 float32) bool { return f1 > f2 }}
	h.Insert([]float32{1.1, 1.2, 1.3, 1.4, 1.5}...)
	actual := h.GetHeap()
	expected := []float32{1.5, 1.4, 1.3, 1.1, 1.2}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestChangePriority1(t *testing.T) {
	h := Heap[int]{}
	h.Comparator = func(i1, i2 int) bool {
		return i1 > i2
	}

	h.Insert([]int{1, 2, 3, 4, 5}...)

	h.ChangePriority(1, 8)
	actual := make([]int, 0, h.Size())
	for h.Size() > 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{8, 5, 3, 2, 1}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}
func TestChangePriority2(t *testing.T) {
	h := Heap[int]{}
	h.Comparator = func(i1, i2 int) bool { return i1 < i2 }

	h.Insert([]int{7, 6, 5, 4, 3, 2, 1}...)

	h.ChangePriority(6, 8)
	actual := make([]int, 0, h.Size())
	for h.Size() > 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{1, 2, 3, 4, 6, 7, 8}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestRemove(t *testing.T) {
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	h.Remove(5)
	h.Remove(2)
	actual := make([]int, 0, h.Size())
	for h.Size() > 0 {
		actual = append(actual, h.ExtractTop())
	}
	expected := []int{9, 8, 5, 4, 3, 2, 1}
	if err := compareArrays(append(actual, expected...)...); err != nil {
		t.Error(err)
	}
}

func TestSizeTopInsert(t *testing.T) {
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 < i2 }}
	h.Insert(9)
	h.Insert(8)
	h.Insert(7)
	h.Insert(1)
	if h.Size() != 4 {
		t.Error("bad size")
	}
	if h.Top() != 1 {
		t.Error("bad top value")
	}
}
