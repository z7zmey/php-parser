package php5_test

import (
	"io/ioutil"
	"testing"

	"github.com/z7zmey/php-parser/internal/php5"
	"github.com/z7zmey/php-parser/internal/scanner"
)

func BenchmarkPhp5(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")
	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		lexer := scanner.NewLexer([]byte(src), "5.6", nil)
		php5parser := php5.NewParser(lexer, nil)
		php5parser.Parse()
	}
}
