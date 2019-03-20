package ast

import (
	"runtime"
	"testing"
)

func BenchmarkEdgeStorageSave(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]Edge, 0, s)
		edgeStore := NewEdgeStorage(buf)

		for i := 0; i < k; i++ {
			e := Edge{
				Type: EdgeTypeStmts,
				From: NodeID(1),
				To:   NodeID(2),
			}
			_ = edgeStore.Create(e)
		}

		if gc {
			runtime.GC()
		}
	}
}
