package sq

import "fmt"

type MedianFinder struct {
	h1 SmallHeap
	h2 BigHeap
}

func (finder *MedianFinder) Init() {
	finder.h1 = SmallHeap{
		Array: make([]int, 0),
	}
	finder.h2 = BigHeap{
		Array: make([]int, 0),
	}
}

func (finder *MedianFinder) AddNum(num int) {
	if len(finder.h1.Array) == len(finder.h2.Array) {
		finder.h2.Insert(num)
		rightBiggest := finder.h2.GetBiggest()
		finder.h2.DeleteBiggest()
		finder.h1.Insert(rightBiggest)
	} else {
		finder.h1.Insert(num)
		leftSmallest := finder.h1.GetSmallest()
		finder.h1.DeleteSmallest()
		finder.h2.Insert(leftSmallest)
	}

	fmt.Println(finder.h1.Array)
	fmt.Println(finder.h2.Array)
}

func (finder *MedianFinder) FindMedian() float64 {
	if len(finder.h1.Array) == len(finder.h2.Array) {
		return float64((finder.h1.Array[0] + finder.h2.Array[0])) / 2
	} else {
		return float64(finder.h1.Array[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
