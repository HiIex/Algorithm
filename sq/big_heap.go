package sq

//"fmt"

// BigHeap 大顶堆
type BigHeap struct {
	Array []int
}

// Insert 插末尾，挨个往上比，直到比不动为止
//  @receiver h
//  @param num
func (h *BigHeap) Insert(num int) {
	h.Array = append(h.Array, num)
	insertPos := len(h.Array) - 1
	parentPos := (insertPos - 1) / 2
	for {
		if parentPos < 0 || h.Array[parentPos] >= h.Array[insertPos] {
			break
		}
		h.Array[insertPos] = h.Array[parentPos]
		h.Array[parentPos] = num
		insertPos = parentPos
		parentPos = (insertPos - 1) / 2
	}

	//fmt.Println(h.Array)
}

func (h BigHeap) GetBiggest() int {
	if len(h.Array) == 0 {
		//fmt.Println("==0")
		return -999
	}

	return h.Array[0]
}

// DeleteSmallest 删第一个，最后一个挪到第一个，挨个往下比直到比不动为止
//  @receiver h
func (h *BigHeap) DeleteBiggest() {
	if len(h.Array) == 0 {
		return
	}

	if len(h.Array) == 1 {
		h.Array = []int{}
		return
	}

	lastNum := h.Array[len(h.Array)-1]
	h.Array[0] = lastNum
	h.Array = h.Array[:len(h.Array)-1]

	insertPos := 0
	leftChildPos := 1
	rightChildPos := 2
	for {
		if leftChildPos >= len(h.Array) {
			break
		}

		var biggerChildPos int
		if rightChildPos >= len(h.Array) {
			biggerChildPos = leftChildPos
		} else {
			if h.Array[leftChildPos] >= h.Array[rightChildPos] {
				biggerChildPos = leftChildPos
			} else {
				biggerChildPos = rightChildPos
			}
		}

		if h.Array[insertPos] >= h.Array[biggerChildPos] {
			break
		}

		h.Array[insertPos] = h.Array[biggerChildPos]
		h.Array[biggerChildPos] = h.Array[insertPos]
		h.Array[biggerChildPos] = lastNum
		insertPos = biggerChildPos
		leftChildPos = 2*insertPos + 1
		rightChildPos = 2*insertPos + 2
	}
}
