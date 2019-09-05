package php7_test

import (
	"io/ioutil"
	"testing"

	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/internal/stxtree"
)

func BenchmarkPhp7(b *testing.B) {
	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser()

	for n := 0; n < b.N; n++ {
		a := &stxtree.Graph{}
		php7parser.Parse([]byte(src), a)
	}
}

func BenchmarkPhp7WithTokens(b *testing.B) {
	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser().WithTokens()

	for n := 0; n < b.N; n++ {
		a := &stxtree.Graph{}
		php7parser.Parse([]byte(src), a)
	}
}

func BenchmarkPhp7Reuse(b *testing.B) {
	src, err := ioutil.ReadFile("bench.php")
	if err != nil {
		b.Error(err)
	}

	php7parser := php7.NewParser()
	a := &stxtree.Graph{}
	for n := 0; n < b.N; n++ {
		a.Reset()
		php7parser.Parse([]byte(src), a)
	}
}
