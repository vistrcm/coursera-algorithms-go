package union_find

import (
	"math/rand"
	"testing"
)

func TestQuickFindUF_Connected(t *testing.T) {
	type fields struct {
		id []int
	}
	type args struct {
		p int
		q int
	}
	type testRun struct {
		connectedTarget args
		want            bool
	}

	tests := []struct {
		name   string
		n      int
		unions []args
		runs   []testRun
	}{
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

func BenchmarkQuickFindUF_Union(b *testing.B) {
	// create uf with b.N comonents
	uf := NewQuickFind(b.N)
	// run union b.N times on random components
	for i := 0; i < b.N; i++ {
		p := rand.Intn(b.N)
		q := rand.Intn(b.N)
		uf.Union(p, q)
	}
}