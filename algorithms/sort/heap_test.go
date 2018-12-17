package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var arr []int
	for i := 0; i < 9; i++ {
		arr = append(arr, r.Intn(100))
	}

	HeapSort(arr, len(arr))
	for i := range arr {
		if i+1 < len(arr) && arr[i] > arr[i+1] {
			t.Error("Failed, HeapSort.")
		}
	}
}
