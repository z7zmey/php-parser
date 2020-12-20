package php7_test

import (
	"io/ioutil"
	"testing"

	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/internal/scanner"
)

func BenchmarkPhp7(b *testing.B) {
	src, err := ioutil.ReadFile("test.php")

	if err != nil {
		b.Fatal("can not read test.php: " + err.Error())
	}

	for n := 0; n < b.N; n++ {
		lexer := scanner.NewLexer(src, "7.4", nil)
		php7parser := php7.NewParser(lexer, nil)
		php7parser.Parse()
	}
}
