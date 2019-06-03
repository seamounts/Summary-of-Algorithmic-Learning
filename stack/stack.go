package stack

type Stack interface {
	Push(string) bool
	Pop() string
	Get() string
	Empty() bool
}

type OrderStack struct {
	items []string
	count int
}

func NewOrderStack() *OrderStack {
	return &OrderStack{
		items: make([]string, 10),
	}
}

func (os *OrderStack) Push(item string) bool {
	if os.count == len(os.items) {
		newItems := make([]string, 0, 2*len(os.items))
		newItems = append(newItems, os.items...)
		os.items = newItems
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

func (os *OrderStack) Get() string {
	if os.count == 0 {
		return ""
	}
	item := os.items[os.count-1]
	return item
}

func (os *OrderStack) Empty() bool {
	return os.count == 0
}
