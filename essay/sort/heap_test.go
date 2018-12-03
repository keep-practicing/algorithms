package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{55, 34, 22, 67, 2, 3, 4, 1}
	HeapSort(arr, 8)
	for i := range arr {
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			t.Error("Failed, HeapSort.")
		}
	}
}
