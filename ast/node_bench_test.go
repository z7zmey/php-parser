package ast

import (
	"runtime"
	"testing"
)

func BenchmarkNodeStorageSave(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]Node, 0, s)
		nodeBuf := NewNodeStorage(buf)

		for i := 0; i < k; i++ {
			n := Node{
				Type: NodeTypeRoot,
			}
			_ = nodeBuf.Create(n)
		}

		if gc {
			runtime.GC()
		}
	}
}
