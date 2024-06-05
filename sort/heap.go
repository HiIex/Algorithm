package sort

func HeapSort(array []int) {
	//构造大顶堆
	for i := len(array)/2 - 1; i >= 0; i-- {
		adjustBigHeap(array, i, len(array)-1)
	}

	//取顶端最大的元素和最后交换，将0 - n-1的堆重新调整，得到一个新的顶端
	for i := len(array) - 1; i >= 0; i-- {
		top := array[0]
		array[0] = array[i]
		array[i] = top
		adjustBigHeap(array, 0, i-1)
	}
}

// adjustBigHeap 调整从top到end的大顶堆
//  @param array
//  @param root
//  @param end
func adjustBigHeap(heap []int, top int, end int) {
	i := top
	temp := heap[top]
	for {
		if 2*i+1 > end {
			break
		}

		var maxIndex int
		if 2*i+1 == end {
			maxIndex = 2*i + 1
		} else if heap[2*i+1] >= heap[2*i+2] {
			maxIndex = 2*i + 1
		} else {
			maxIndex = 2*i + 2
		}

		//用temp比
		if temp < heap[maxIndex] {
			heap[i] = heap[maxIndex]
			i = maxIndex
		} else {
			break
		}
	}

	//容易错的地方
	heap[i] = temp
}
