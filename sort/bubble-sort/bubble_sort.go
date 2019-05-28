package bubble_sort

func BubbleSort(datas []int) []int {
	for i := 0; i < len(datas)-1; i++ {
		for j := 0; j < len(datas)-1-i; j++ {
			if datas[j] > datas[j+1] {
				tmp := datas[j]
				datas[j] = datas[j+1]
				datas[j+1] = tmp
			}
		}

	}
	return datas
}
