package sort

func Merge(array []int,start int,end int) {
	if start>=end{
		return 
	}

	mid:=(start+end)/2
	Merge(array,start,mid)
	Merge(array,mid+1,end)

	tempArray:=make([]int,0)
	for k:=start;k<=end;k++{
		tempArray=append(tempArray, array[k])
	}

	i:=0
	j:=mid+1-start
	for k:=start;k<=end;k++{
		if i==mid-start+1{
			array[k]=tempArray[j]
			j++
			continue
		}

		if j==end-start+1{
			array[k]=tempArray[i]
			i++
			continue
		}

		//容易错：比的是tempArray
		if tempArray[i]<=tempArray[j]{
			array[k]=tempArray[i]
			i++
		}else{
			array[k]=tempArray[j]
			j++
		}
	}
}

/* func Merge(array []int, start int, end int) {
	if start >= end {
		return 
	}

	mid := (start + end) / 2
	Merge(array, start, mid)
	Merge(array, mid+1, end)

	tempArray := make([]int, 0)
	for i := start; i <= end; i++ {
		tempArray=append(tempArray, array[i])
	}

	i := 0
	j := mid -start+1
	for k := start; k <= end; k++ {
		if i == mid-start+1 {
			array[k] = tempArray[j]
			j++
			continue
		}

		if j == end-start+1 {
			array[k] = tempArray[i]
			i++
			continue
		}

		if tempArray[i] <= tempArray[j] {
			array[k] = tempArray[i]
			i++
		} else {
			array[k] = tempArray[j]
			j++
		}
	}

} */
