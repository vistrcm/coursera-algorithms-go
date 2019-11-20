package union_find

//QuickUnionUF implements Weighted Quick Union algorithm with path compression
type QuickUnionWPCUF struct {
	id []int
	sz []int
}

//NewQuickUnionWPCUF initializes a new QuickFindUF instance
func NewQuickUnionWPCUF(N int) QuickUnionWPCUF {
	id := make([]int, N)
	sz := make([]int, N)
	for pos := range id {
		id[pos] = pos
	}

	return QuickUnionWPCUF{id: id, sz: sz}
}

//root finds the root of element
func (uf *QuickUnionWPCUF) root(i int) int {
	for i != uf.id[i] {
		uf.id[i] = uf.id[uf.id[i]]
		i = uf.id[i]
	}
	return i
}

//Connected checks if p and q are connected.
func (uf *QuickUnionWPCUF) Connected(p, q int) bool {
	return uf.root(p) == uf.root(q)
}

//Union connects p and q
func (uf *QuickUnionWPCUF) Union(p, q int) {
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
