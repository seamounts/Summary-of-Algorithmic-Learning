package insert_sort

func InsertSort(datas []int) []int {
	for i := 1; i < len(datas); i++ {
		if datas[i] >= datas[i-1] {
			continue
		}
		for j := 0; j < i; j++ {
			if datas[i] < datas[j] {
				tmp := datas[i]
				k := i
				for {
					if k == j {
						break
					}
					datas[k] = datas[k-1]
					k = k - 1
				}
				datas[j] = tmp
				break
			}
		}
	}
	return datas
}
