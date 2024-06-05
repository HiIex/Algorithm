package array

/* 设备中存有 n 个文件，文件 id 记于数组 documents。
若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。
注意！！！：0 ≤ documents[i] ≤ n-1

输入：documents = [2, 5, 3, 0, 5, 0]
输出：0 或 5
*/

// FindRepeatDocument 把数组本身当哈希表，按索引与值一一对应，若出现重复必会出现documents[tbd]==documents[documents[tbd]]
//  @param documents
//  @return int
func FindRepeatDocument(documents []int) int {
	tbd := 0
	for {
		//若刚好没有重复，tbd会滑到len(documents)
		//若最后才出现重复，则当tbd==len(documents)-1时会判断出
		if tbd == len(documents) {
			break
		}

		//若值已经和索引相等，则该元素已在对应位置，取++tbd的位置接着比
		if documents[tbd] == tbd {
			tbd++
			continue
		}

		if documents[tbd] == documents[documents[tbd]] {
			return documents[tbd]
		}

		temp := documents[tbd]
		documents[tbd] = documents[temp]
		documents[temp] = temp
	}

	return -1
}
