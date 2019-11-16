package union_find

import (
	"math/rand"
	"testing"
)

// shared test data
type args struct {
	p int
	q int
}
type testRun struct {
	connectedTarget args
	want            bool
}

type testCases struct {
	name   string
	n      int
	unions []args
	runs   []testRun
}

var (
	tests = []testCases{
		{
			name:   "2",
			n:      2,
			unions: []args{},
			runs: []testRun{
				{
					connectedTarget: args{0, 1},
					want:            false,
				},
			},
		},

		{
			name: "10",
			n:    10,
			unions: []args{
				{4, 3},
				{3, 8},
				{6, 5},
				{9, 4},
				{2, 1},
				{8, 9},
				{5, 0},
				{7, 2},
				{6, 1},
			},
			runs: []testRun{
				{
					connectedTarget: args{0, 1}, want: true,
				},
				{
					connectedTarget: args{0, 2}, want: true,
				},
				{
					connectedTarget: args{0, 3}, want: false,
				},
				{
					connectedTarget: args{0, 4}, want: false,
				},
				{
					connectedTarget: args{0, 5}, want: true,
				},
				{
					connectedTarget: args{0, 6}, want: true,
				},
				{
					connectedTarget: args{0, 7}, want: true,
				},
				{
					connectedTarget: args{0, 8}, want: false,
				},
				{
					connectedTarget: args{0, 9}, want: false,
				},
				{
					connectedTarget: args{6, 8}, want: false,
				},
				{
					connectedTarget: args{8, 9}, want: true,
				},
			},
		},
	}
)

func TestQuickFindUF_Connected(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewQuickFind(tt.n)
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

func benchmarkQuickFindUF_Union(size int, b *testing.B) {
	// create uf with b.N comonents
	uf := NewQuickFind(size)
	// run union b.N times on random components
	for i := 0; i < b.N; i++ {
		p := rand.Intn(size)
		q := rand.Intn(size)
		uf.Union(p, q)
	}
}

// benchmark QuickFindUF_Union with different inputs

func BenchmarkQuickFindUF_Union1(b *testing.B)      { benchmarkQuickFindUF_Union(1, b) }
func BenchmarkQuickFindUF_Union10(b *testing.B)     { benchmarkQuickFindUF_Union(10, b) }
func BenchmarkQuickFindUF_Union100(b *testing.B)    { benchmarkQuickFindUF_Union(100, b) }
func BenchmarkQuickFindUF_Union1000(b *testing.B)   { benchmarkQuickFindUF_Union(1000, b) }
func BenchmarkQuickFindUF_Union10000(b *testing.B)  { benchmarkQuickFindUF_Union(10000, b) }
func BenchmarkQuickFindUF_Union100000(b *testing.B) { benchmarkQuickFindUF_Union(100000, b) }

func TestQuickUnionUF_Connected(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uf := NewQuickUnionUF(tt.n)
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

func benchmarkQuickUnionUF_Union(size int, b *testing.B) {
	// create uf with b.N comonents
	uf := NewQuickUnionUF(size)
	// run union b.N times on random components
	for i := 0; i < b.N; i++ {
		p := rand.Intn(size)
		q := rand.Intn(size)
		uf.Union(p, q)
	}
}

// benchmark QuickFindUF_Union with different inputs

func BenchmarkQuickUnionUF_Union1(b *testing.B)      { benchmarkQuickUnionUF_Union(1, b) }
func BenchmarkQuickUnionUF_Union10(b *testing.B)     { benchmarkQuickUnionUF_Union(10, b) }
func BenchmarkQuickUnionUF_Union100(b *testing.B)    { benchmarkQuickUnionUF_Union(100, b) }
func BenchmarkQuickUnionUF_Union1000(b *testing.B)   { benchmarkQuickUnionUF_Union(1000, b) }
func BenchmarkQuickUnionUF_Union10000(b *testing.B)  { benchmarkQuickUnionUF_Union(10000, b) }
func BenchmarkQuickUnionUF_Union100000(b *testing.B) { benchmarkQuickUnionUF_Union(100000, b) }
