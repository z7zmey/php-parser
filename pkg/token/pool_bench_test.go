package token

import (
	"testing"
)

const amount = 100000

func BenchmarkPlain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)

		for i := 0; i < amount; i++ {
			buf = append(buf, &Token{})
		}
	}
}

func BenchmarkSlice128(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 128)

		for i := 0; i < amount; i++ {
			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 512)

		for i := 0; i < amount; i++ {
			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice1024(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 1024)

		for i := 0; i < amount; i++ {
			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkSlice2048(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 2048)

		for i := 0; i < amount; i++ {
			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend128(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 128)

		for i := 0; i < amount; i++ {
			if len(slc) == 128 {
				slc = make([]Token, 0, 128)
			}

			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 512)

		for i := 0; i < amount; i++ {
			if len(slc) == 512 {
				slc = make([]Token, 0, 512)
			}

			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend1024(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 1024)

		for i := 0; i < amount; i++ {
			if len(slc) == 1024 {
				slc = make([]Token, 0, 1024)
			}

			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkBlockAppend2048(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := make([]*Token, 0, amount)
		slc := make([]Token, 0, 2048)

		for i := 0; i < amount; i++ {
			if len(slc) == 2048 {
				slc = make([]Token, 0, 2048)
			}

			slc = append(slc, Token{})
			buf = append(buf, &slc[len(slc)-1])
		}
	}
}

func BenchmarkPool128(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pool := NewPool(128)
		buf := make([]*Token, 0, amount)

		for i := 0; i < amount; i++ {
			buf = append(buf, pool.Get())
		}
	}
}

func BenchmarkPool512(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pool := NewPool(512)
		buf := make([]*Token, 0, amount)

		for i := 0; i < amount; i++ {
			buf = append(buf, pool.Get())
		}
	}
}

func BenchmarkPool1024(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pool := NewPool(1024)
		buf := make([]*Token, 0, amount)

		for i := 0; i < amount; i++ {
			buf = append(buf, pool.Get())
		}
	}
}

func BenchmarkPool2048(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pool := NewPool(2048)
		buf := make([]*Token, 0, amount)

		for i := 0; i < amount; i++ {
			buf = append(buf, pool.Get())
		}
	}
}
