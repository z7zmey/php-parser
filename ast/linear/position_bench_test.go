package linear

import (
	"runtime"
	"testing"

	"github.com/z7zmey/php-parser/ast"
)

const k = 3000
const s = 1024
const gc = false

func BenchmarkPositionStorageSave(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]ast.Position, 0, s)
		posBuf := NewPositionStorage(buf)

		for i := 0; i < k; i++ {
			p := ast.Position{1, 1, 1, 1}
			_ = posBuf.Create(p)
		}

		if gc {
			runtime.GC()
		}
	}
}
