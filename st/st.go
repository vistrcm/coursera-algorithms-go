package st

type Comparable interface {
	//CompareTo should return negative int if other is bigger, 0 if equal and positive int if less
	CompareTo(other Comparable) int
}

type Key Comparable
type Value string

type Node struct {
	key   Key
	value Value
	left  *Node
	right *Node
}

func NewNode(key Key, value Value) *Node {
	return &Node{key: key, value: value}
}

type BST struct {
	root *Node

}

func (bst BST)Put(key Key, value Value){
	panic("NOT implemented")
}

func (bst BST)Get(key Key) Value{
	x := bst.root
	for x!=nil {
		cmp := key.CompareTo(x.key)
		if cmp <0 {x = x.left}
		else if cmp > 0 {x}
	}
}

func (bst BST)Delete(key Key){
	panic("NOT implemented")
}

func (bst BST)Keys() []Key{
	panic("NOT implemented")
}