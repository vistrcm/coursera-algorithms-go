package rbbst

import "github.com/vistrcm/coursera-algorithms-pi-go/st"

//Left-leaning Red-Black BST

const (
	RED   = true
	BLACK = false
)

type color bool

type node struct {
	key   st.Key
	value st.Value
	left  *node
	right *node
	count int
	color color // color of link to the node
}

func newNode(key st.Key, value st.Value, color color) *node {
	return &node{key: key, value: value, color: color}
}

type RBBST struct {
	root *node
}

//Get the value for the key
func (t RBBST) Get(key st.Key) st.Value {
	x := t.root
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

func (t *RBBST) Put(key st.Key, value st.Value) {
	t.root = put(t.root, key, value)
}

func put(h *node, key st.Key, value st.Value) *node {
	if h == nil {
		return newNode(key, value, RED)
	}
	cmp := key.CompareTo(h.key)
	if cmp < 0 {
		h.left = put(h.left, key, value)
	} else if cmp > 0 {
		h.right = put(h.right, key, value)
	} else {
		h.value = value
	}

	// rotations and flips
	// lean left
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	// balance 4-node
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	//split 4-node
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	h.count = 1 + size(h.left) + size(h.right)

	return h
}

//isRed test if node is red
func isRed(x *node) bool {
	if x == nil {
		return false
	}
	return x.color == RED //nolint:gosimple //it is not easier to read x.color instead of comparison
}

func rotateLeft(h *node) *node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

func rotateRight(h *node) *node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	return x
}

func flipColors(h *node) {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
}

func size(x *node) int {
	if x == nil {
		return 0
	}
	return x.count
}

func (bst *RBBST) Keys() []st.Key {
	var q []st.Key
	inorder(bst.root, &q)
	return q
}

func inorder(x *node, q *[]st.Key) {
	if x == nil {
		return
	}
	inorder(x.left, q)
	*q = append(*q, x.key)
	inorder(x.right, q)
}
