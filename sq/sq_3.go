package sq

/*
输入：putIn = [6,7,8,9,10,11], takeOut = [9,11,10,8,7,6]
输出：true
解释：我们可以按以下操作放入并拿取书籍：
push(6), push(7), push(8), push(9), pop() -> 9,
push(10), push(11),pop() -> 11,pop() -> 10, pop() -> 8, pop() -> 7, pop() -> 6

输入：putIn = [6,7,8,9,10,11], takeOut = [11,9,8,10,6,7]
输出：false
解释：6 不能在 7 之前取出。
*/

// ValidateBookSequences 用一个栈存储，能pop出去就pop，不能就接着push，没得push也不能pop就表示为false
//  @param putIn
//  @param takeOut
//  @return bool
func ValidateBookSequences(putIn []int, takeOut []int) bool {
	if len(putIn) == 0 {
		return true
	}

	s := make([]int, 0)
	i := 0
	j := 0
	length := len(putIn)
	s = append(s, putIn[0])

	for {
		if j == length {
			return true
		}

		top := len(s) - 1
		if top >= 0 && takeOut[j] == s[top] {
			j++
			s = s[:top]
			continue
		}

		if i == length-1 {
			return false
		} else {
			i++
			s = append(s, putIn[i])
		}

	}
}
