package array

/*
每部分文件编号均为一个 正整数（至少含有两个文件）。传输要求为：连续文件编号总和为接收方指定数字 target 的所有文件。请返回所有符合该要求的文件传输组合列表。
每种组合按照文件编号 升序 排列；
不同组合按照第一个文件编号 升序 排列。

输入：target = 18
输出：[[3,4,5,6],[5,6,7]]
*/

// FileCombination 滑动窗口
//  @param target
//  @return []
func FileCombination(target int) [][]int {
	var result [][]int
	start := 1
	end := 1

	for {
		//头尾相加求和
		sum := (start + end) * (end - start + 1) / 2

		//达到这个情况肯定没后续了
		if 2*start >= target {
			break
		}

		if sum == target {
			array := make([]int, end-start+1)
			for i := 0; i <= end-start; i++ {
				array[i] = start + i
			}
			result = append(result, array)

			//等于目标start往后滑一位，寻找新组合
			start++
			continue
		}

		//小于目标end往后滑一位
		if sum < target {
			end++
			continue
		}

		//大于目标start往后滑一位
		if sum > target {
			start++
			continue
		}
	}

	return result
}
