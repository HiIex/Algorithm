package sort

func Fast(array []int,start int, end int ) []int{
	if start>=end{
		return array
	}

	pivot:=partition(array,start,end)
	Fast(array,start,pivot)
	Fast(array,pivot+1,end)
	return array
}

func partition(array []int, start int ,end int ) int{
	i:=start
	j:=end
	pivot:=start
	for{
		if i>=j{
			break
		}

		for {
			if i<j&&array[j]>=array[pivot]{
				j--
			}else{
				break
			}
		}

		for {
			if i<j&&array[i]<=array[pivot]{
				i++
			}else{
				break
			}
		}

		swap(array,i,j)
	}

	swap(array,pivot,i)
	return i
}

/* func Fast(array []int, start int, end int) []int {
	if start >= end {
		return array
	}

	pivot := partition(array, start, end)
	Fast(array, start, pivot-1)
	Fast(array, pivot+1, end)
	return array
}

func partition(array []int, start int, end int) int {
	i := start
	j := end
	key := start
	for {
		if i >= j {
			break
		}

		for {
			if i < j && array[j] >= array[key] {
				j--
			} else {
				break
			}
		}

		for {
			if i < j && array[i] <= array[key] {
				i++
			} else {
				break
			}

		}

		swap(array, i, j)
	}

	swap(array, start, i)
	return i
} */

func swap(array []int, a int, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}
