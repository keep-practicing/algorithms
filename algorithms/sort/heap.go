package sort

func shiftDown(arr []int, n int, k int) {
	for 2*k+1 < n {
		j := 2*k + 1
		if j+1 < n && arr[j+1] > arr[j] {
			j++
		}
		if arr[k] >= arr[j] {
			break
		}
		arr[j], arr[k] = arr[k], arr[j]
		k = j
	}
}

// HeapSort 堆排序
func HeapSort(arr []int, n int) {
	// heapify
	for i := (n - 1) / 2; i >= 0; i-- {
		shiftDown(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		shiftDown(arr, i, 0)
	}
}
