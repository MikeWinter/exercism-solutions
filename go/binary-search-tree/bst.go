package binarysearchtree

type SearchTreeData struct {
	data        int
	left, right *SearchTreeData
}

func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

func (t *SearchTreeData) Insert(value int) {
	var insertionPoint **SearchTreeData
	node := t
	newNode := &SearchTreeData{data: value}
	for {
		if value <= node.data {
			if node.left == nil {
				insertionPoint = &node.left
				break
			}
			node = node.left
		} else {
			if node.right == nil {
				insertionPoint = &node.right
				break
			}
			node = node.right
		}
	}
	*insertionPoint = newNode
}

func (t *SearchTreeData) MapString(mapper func(int) string) []string {
	strings := make([]string, 0)
	t.inOrder(func(value int) {
		strings = append(strings, mapper(value))
	})
	return strings
}

func (t *SearchTreeData) MapInt(mapper func(int) int) []int {
	ints := make([]int, 0)
	t.inOrder(func(value int) {
		ints = append(ints, mapper(value))
	})
	return ints
}

func (t *SearchTreeData) inOrder(callback func(int)) {
	if t.left != nil {
		t.left.inOrder(callback)
	}

	callback(t.data)

	if t.right != nil {
		t.right.inOrder(callback)
	}
}
