package parser

import (
	"github.com/z7zmey/php-parser/internal/php5"
	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/internal/version"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() ast.Vertex
}

type Config struct {
	WithTokens       bool
	WithPositions    bool
	ErrorHandlerFunc func(e *errors.Error)
}

func Parse(src []byte, ver string, cfg Config) (ast.Vertex, error) {
	var parser Parser

	r, err := version.Compare(ver, "7.0")
	if err != nil {
		return nil, err
	}

	lexer := scanner.NewLexer(src, ver, cfg.WithTokens, cfg.ErrorHandlerFunc)

	if r == -1 {
		parser = php5.NewParser(lexer, cfg.ErrorHandlerFunc)
	} else {
		parser = php7.NewParser(lexer, cfg.ErrorHandlerFunc)
	}

	parser.Parse()

	return parser.GetRootNode(), nil
}
