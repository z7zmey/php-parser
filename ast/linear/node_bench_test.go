package linear

import (
	"runtime"
	"testing"

	"github.com/z7zmey/php-parser/ast"
)

func BenchmarkNodeStorageSave(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]Node, 0, s)
		nodeBuf := NewNodeStorage(buf)

		for i := 0; i < k; i++ {
			n := Node{
				Type: ast.NodeTypeRoot,
			}
			_ = nodeBuf.Create(n)
		}

		if gc {
			runtime.GC()
		}
	}
}
