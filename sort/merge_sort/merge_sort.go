package merge_sort

func MergeSort(datas []int) []int {
	if len(datas) == 1 {
		return datas
	}
	left := MergeSort(datas[0 : len(datas)/2])
	right := MergeSort(datas[len(datas)/2 : len(datas)])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	sortDatas := make([]int, 0, leftLen+rightLen)

	i := 0
	j := 0

	for {
		if i == leftLen && j == rightLen {
			break
		}
		if i == leftLen {
			sortDatas = append(sortDatas, right[j:rightLen]...)
			break
		}
		if j == rightLen {
			sortDatas = append(sortDatas, left[i:leftLen]...)
			break
		}

		if left[i] < right[j] {
			sortDatas = append(sortDatas, left[i])
			i++
		} else {
			sortDatas = append(sortDatas, right[j])
			j++
		}
	}
	return sortDatas
}
