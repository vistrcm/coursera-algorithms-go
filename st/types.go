package st

type Comparable interface {
	//CompareTo should return negative int if other is bigger, 0 if equal and positive int if less
	CompareTo(other interface{}) int
}

type Key Comparable
type Value interface{}
