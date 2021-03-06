package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickUnionWPCUF_Connected(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewQuickUnionWPCUF(tt.n)
			for _, union := range tt.unions {
				uf.Union(union.p, union.q)
			}

			for _, run := range tt.runs {
				p := run.connectedTarget.p
				q := run.connectedTarget.q
				if got := uf.Connected(p, q); got != run.want {
					t.Errorf("Connected(%d, %d) = %v, want %v", p, q, got, run.want)
				}
			}

		})
	}
}

func benchmarkQuickUnionWPCUF_Union(size int, b *testing.B) {
	// create uf with b.N components
	uf := NewQuickUnionWPCUF(size)
	// run union b.N times on random components
	for i := 0; i < b.N; i++ {
		p := rand.Intn(size)
		q := rand.Intn(size)
		uf.Union(p, q)
	}
}

// benchmark QuickFindUF_Union with different inputs

func BenchmarkQuickUnionWPCUF_Union1(b *testing.B)     { benchmarkQuickUnionWPCUF_Union(1, b) }
func BenchmarkQuickUnionWPCUF_Union10(b *testing.B)    { benchmarkQuickUnionWPCUF_Union(10, b) }
func BenchmarkQuickUnionWPCUF_Union100(b *testing.B)   { benchmarkQuickUnionWPCUF_Union(100, b) }
func BenchmarkQuickUnionWPCUF_Union1000(b *testing.B)  { benchmarkQuickUnionWPCUF_Union(1000, b) }
func BenchmarkQuickUnionWPCUF_Union10000(b *testing.B) { benchmarkQuickUnionWPCUF_Union(10000, b) }
