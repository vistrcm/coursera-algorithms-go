package st

type node struct {
	key   Key
	value Value
	left  *node
	right *node
	count int
}

func newNode(key Key, value Value) *node {
	return &node{key: key, value: value}
}

type BST struct {
	root *node
}

func (bst *BST) Put(key Key, value Value) {
	bst.root = put(bst.root, key, value)
}

func put(x *node, key Key, value Value) *node {
	if x == nil {
		return newNode(key, value)
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = put(x.left, key, value)
	} else if cmp > 0 {
		x.right = put(x.right, key, value)
	} else {
		x.value = value
	}
	x.count = 1 + size(x.left) + size(x.right)
	return x
}

func (bst BST) Get(key Key) Value {
	x := bst.root
	for x != nil {
		cmp := key.CompareTo(x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.value
		}
	}
	return nil
}

func (bst *BST) Delete(key Key) {
	bst.root = del(bst.root, key)
}

func del(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = del(x.left, key)
	} else if cmp > 0 {
		x.right = del(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}

		t := x
		x = min(t.right)
		x.right = deleteMin(t.right)
		x.left = t.left
	}
	x.count = size(x.left) + size(x.right) + 1
	return x
}

func (bst *BST) Min() Key {
	return min(bst.root).key
}

func min(x *node) *node {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func (bst *BST) Keys() []Key {
	var q []Key
	inorder(bst.root, &q)
	return q
}

func inorder(x *node, q *[]Key) {
	if x == nil {
		return
	}
	inorder(x.left, q)
	*q = append(*q, x.key)
	inorder(x.right, q)
}

func (bst BST) Floor(key Key) Key {
	x := floor(bst.root, key)
	if x == nil {
		return nil
	}
	return x.key
}

func floor(x *node, key Key) *node {
	if x == nil {
		return nil
	}

	cmp := key.CompareTo(x.key)

	if cmp == 0 {
		return x
	}

	if cmp < 0 {
		return floor(x.left, key)
	}

	t := floor(x.right, key)
	if t != nil {
		return t
	} else {
		return x
	}
}

func (bst BST) Size() int {
	return size(bst.root)
}

func size(x *node) int {
	if x == nil {
		return 0
	}
	return x.count
}

func (bst BST) Rank(key Key) int {
	return rank(key, bst.root)
}

func rank(key Key, x *node) int {
	if x == nil {
		return 0
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		return rank(key, x.left)
	} else if cmp > 0 {
		return 1 + size(x.left) + rank(key, x.right)
	} else {
		return size(x.left)
	}
}

func (bst *BST) DeleteMin() {
	bst.root = deleteMin(bst.root)
}

func deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = deleteMin(x.left)
	x.count = 1 + size(x.left) + size(x.right)
	return x
}
