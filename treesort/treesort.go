// 二叉树实现插入排序
package treesort

type tree struct {
	value       int
	left, right *tree
}

// Sort 就地排序
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues 将元素按顺序追加到 values 里面， 然后返回结果slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// 等价于返回&tree{value: value}
		// new函数创建一个指定类型的结构体并返回其指针
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
