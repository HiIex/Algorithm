package array

/* 给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。

螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。

输入：array = [[1,2,3],[8,9,4],[7,6,5]]
输出：[1,2,3,4,5,6,7,8,9]

*/

// SpiralArray 维护上下左右四个边缘的值
//  @param array 
//  @return []int 
func SpiralArray(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	left := 0
	top := 0
	right := len(array[0]) - 1
	bottom := len(array) - 1
	result := make([]int, 0)

	for {
		for j := left; j <= right; j++ {
			result = append(result, array[top][j])
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			result = append(result, array[i][right])
		}
		right--
		if right < left {
			break
		}

		for j := right; j >= left; j-- {
			result = append(result, array[bottom][j])
		}
		bottom--
		if bottom < top {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, array[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}
