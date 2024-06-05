package main

import (
	//"algorithm/sort"
	"fmt"
)

var (
	visited [][]bool
)

func wardrobeFinishing(m int, n int, cnt int) int {
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			visited[i][j] = false
		}
	}
	return dfs(0, 0, 0, 0, m, n, cnt)

}

func dfs(i int, j int, di int, dj int, m int, n int, cnt int) int {
	if i == m {
		return 0
	}

	if j == n {
		return 0
	}

	if visited[i][j] {
		return 0
	}

	if di+dj > cnt {
		return 0
	}

	visited[i][j] = true
	fmt.Println(i,j)

	var newDI int
	if (di+1)%10 != 0 {
		newDI = di + 1
	} else {
		newDI = di - 8
	}
	res1 := dfs(i+1, j, newDI, dj, m, n, cnt)

	var newDJ int
	if (dj+1)%10 != 0 {
		newDJ = dj + 1
	} else {
		newDJ = dj - 8
	}
	res2 := dfs(i, j+1, di, newDJ, m, n, cnt)

	return 1 + res1 + res2
}

func main() {
	//fmt.Print(sq.MaxAltitude([]int{14,2,27,-5,28,13,39},3))

	/* 	h := sq.BigHeap{
	   		Array: make([]int, 0),
	   	}
	   	h.Insert(5)
	   	h.Insert(-1)
	   	h.Insert(-2)
	   	fmt.Println(h.GetBiggest())
	   	h.DeleteSmallest()
	   	fmt.Println(h.GetBiggest())
	   	h.Insert(3)
	   	h.Insert(4)
	   	h.DeleteSmallest()
	   	fmt.Println(h.GetBiggest())
	   	h.Insert(-5)
	   	h.Insert(6)
	   	fmt.Println(h.GetBiggest())
	   	h.DeleteSmallest()
	   	fmt.Println(h.GetBiggest()) */

	/* 	finder:=sq.MedianFinder{}
	   	finder.Init()
	   	finder.AddNum(3)
	   	fmt.Println(finder.FindMedian())
	   	finder.AddNum(5)
	   	fmt.Println(finder.FindMedian())
	   	finder.AddNum(6)
	   	fmt.Println(finder.FindMedian())
	   	finder.AddNum(-1)
	   	fmt.Println(finder.FindMedian())
	   	finder.AddNum(-5)
	   	fmt.Println(finder.FindMedian())
	   	finder.AddNum(3)
	   	fmt.Println(finder.FindMedian()) */

	//a := []int{5, 6, -1, 4, 5, 1, 3, 2, 6, 8, 9, -5}

	//fmt.Println(sort.Bubble(a))
	//fmt.Println(sort.Fast(a,0,len(a)-1))
	//sort.Merge(a,0,len(a)-1)
	//sort.HeapSort(a)
	//fmt.Println(a)
	fmt.Println(wardrobeFinishing(20, 20, 9))
}
