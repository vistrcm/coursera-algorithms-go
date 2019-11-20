package union_find

//QuickUnionUF implements Quick Union algorithm
type QuickUnionWUF struct {
	id []int
	sz []int
}

//NewQuickUnionWUF initializes a new QuickFindUF instance
func NewQuickUnionWUF(N int) QuickUnionWUF {
	id := make([]int, N)
	sz := make([]int, N)
	for pos, _ := range id {
		id[pos] = pos
	}

	return QuickUnionWUF{id: id, sz: sz}
}

//root finds the root of element
func (uf *QuickUnionWUF) root(i int) int {
	for i != uf.id[i] {
		i = uf.id[i]
	}
	return i
}

//Connected checks if p and q are connected.
func (uf *QuickUnionWUF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

//Union connects p and q
func (uf *QuickUnionWUF) Union(p, q int) {
	pRoot := uf.root(p)
	qRoot := uf.root(q)
	if pRoot == qRoot {
		return
	}

	if uf.sz[pRoot] < uf.sz[qRoot] {
		uf.id[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	} else {
		uf.id[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}

}
