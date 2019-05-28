package select_sort

func SelectSort(datas []int) []int {

	for i := 0; i < len(datas)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(datas); j++ {
			if datas[minIndex] > datas[j] {
				minIndex = j
			}
		}
		tmp := datas[i]
		datas[i] = datas[minIndex]
		datas[minIndex] = tmp
	}
	return datas
}
