package quick_sort

func QuickSort(datas []int) []int {
	QuickSort_C(datas, 0, len(datas)-1)
	return datas
}

func QuickSort_C(datas []int, p, r int) {
	if p >= r {
		return
	}
	pivot := datas[r]
	i := p
	for j := p; j < r; j++ {
		if datas[j] < pivot {
			tmp := datas[j]
			datas[j] = datas[i]
			datas[i] = tmp
			i++
		}
	}
	tmp := datas[i]
	datas[i] = datas[r]
	datas[r] = tmp

	QuickSort_C(datas, 0, i-1)
	QuickSort_C(datas, i+1, r)
}
