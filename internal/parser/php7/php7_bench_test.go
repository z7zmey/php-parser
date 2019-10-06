package php7_test

import (
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/internal/tree"
)

func BenchmarkPhp7(b *testing.B) {
	ballast := make([]byte, 100<<20)

	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser()

	for n := 0; n < b.N; n++ {
		a := tree.NewTree(1024)
		php7parser.Parse([]byte(src), a)
	}

	runtime.KeepAlive(ballast)
}

func BenchmarkPhp7WithTokens(b *testing.B) {
	ballast := make([]byte, 100<<20)

	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser().WithTokens()

	for n := 0; n < b.N; n++ {
		a := tree.NewTree(1024)
		php7parser.Parse([]byte(src), a)
	}

	runtime.KeepAlive(ballast)
}

func BenchmarkPhp7Reuse(b *testing.B) {
	ballast := make([]byte, 100<<20)

	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser()
	a := tree.NewTree(1024)
	for n := 0; n < b.N; n++ {
		a.Reset()
		php7parser.Parse([]byte(src), a)
	}

	runtime.KeepAlive(ballast)
}
