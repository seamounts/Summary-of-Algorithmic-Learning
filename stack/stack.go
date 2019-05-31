package stack

type Stack interface {
	Push(string) bool
	Pop() string
}

type OrderStack struct {
	items [10]string
	count int
}

func (os *OrderStack) Push(item string) bool {
	if os.count == len(os.items) {
		return false
	}
	os.items[os.count] = item
	os.count++
	return true
}

func (os *OrderStack) Pop() string {
	if os.count == 0 {
		return ""
	}
	item := os.items[os.count-1]
	os.count--
	return item
}
