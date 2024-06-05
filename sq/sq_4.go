package sq

import (
	"fmt"
)

/*
输入：heights = [14,2,27,-5,28,13,39], limit = 3
输出：[27,27,28,28,39]
解释：
  滑动窗口的位置                最大值
---------------               -----
[14 2 27] -5 28 13 39          27
14 [2 27 -5] 28 13 39          27
14 2 [27 -5 28] 13 39          28
14 2 27 [-5 28 13] 39          28
14 2 27 -5 [28 13 39]          39
*/

// MaxAltitude 降序队列，对头最大，队尾最小，始终维持从大到小
// 进队：如果队列中的元素比新来的元素小就删掉，因为肯定用不上了
// 出队：如果出队的元素的值合降序队列中的队头值一样，降序队列中的队头出队（只可能在队头，或者不在降序队列里）（假设不在降序队头，说明之后进队的元素有比当前队头元素大的，那么该元素进队的时候会导致当前元素在降序队列中被删）
//
//	@param heights
//	@param limit
//	@return []int
func MaxAltitude(heights []int, limit int) []int {
	q := make([]int, 0)
	q = append(q, heights[0])
	//建立长度为limit的滑动窗口
	for i := 1; i < limit; i++ {
		tail := len(q) - 1
		for {
			if q[tail] >= heights[i] {
				break
			}

			q = q[:tail]
			tail--
			if tail < 0 {
				break
			}
		}
		q = append(q, heights[i])
		fmt.Println(q)
	}

	//滑动
	highest := []int{q[0]}
	for j := limit; j < len(heights); j++ {
		if q[0] == heights[j-limit] {
			q = q[1:]
		}

		tail := len(q) - 1
		if tail < 0 {
			q = append(q, heights[j])
			highest = append(highest, heights[j])
			continue
		}

		for {
			if q[tail] >= heights[j] {
				break
			}
			q = q[:tail]
			tail--
			if tail < 0 {
				break
			}
		}

		q = append(q, heights[j])
		highest = append(highest, q[0])
	}

	return highest
}
