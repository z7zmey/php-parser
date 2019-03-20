package ast

import (
	"runtime"
	"testing"
)

const k = 3000
const s = 1024
const gc = false

func BenchmarkPositionStorageSave(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]Position, 0, s)
		posBuf := NewPositionStorage(buf)

		for i := 0; i < k; i++ {
			p := Position{1, 1, 1, 1}
			_ = posBuf.Create(p)
		}

		if gc {
			runtime.GC()
		}
	}
}
