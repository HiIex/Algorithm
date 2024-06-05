package sort

func Bubble(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}

	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}

	return a
}
