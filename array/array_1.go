package array

// TrainingPlan 将所有奇数调整至偶数之前,不用考虑顺序（奇数在前，偶数在后）
//  @param actions
//  @return []int
func TrainingPlan(actions []int) []int {
	//找第一个偶数的位置
	firstEvenIndex := -1
	for i := 0; i < len(actions); i++ {
		if actions[i]%2 == 0 {
			firstEvenIndex = i
			break
		}
	}

	if firstEvenIndex == -1 {
		return actions
	}

	//从第一个偶数的位置接着遍历，之后出现奇数便和第一个偶数的位置互换，第一个偶数的index+1
	for j := firstEvenIndex + 1; j < len(actions); j++ {
		if actions[j]%2 == 1 {
			temp := actions[firstEvenIndex]
			actions[firstEvenIndex] = actions[j]
			actions[j] = temp
			firstEvenIndex++
		}
	}

	return actions
}
