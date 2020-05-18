package parser

import (
	"github.com/z7zmey/php-parser/internal/php5"
	"github.com/z7zmey/php-parser/internal/php7"
	"github.com/z7zmey/php-parser/internal/version"
	"github.com/z7zmey/php-parser/pkg/ast"
	"github.com/z7zmey/php-parser/pkg/errors"
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() ast.Vertex
	GetErrors() []*errors.Error
}

func NewParser(src []byte, v string, withTokens bool) (Parser, error) {
	var parser Parser

	r, err := version.Compare(v, "7.0")
	if err != nil {
		return nil, err
	}

	if r == -1 {
		parser = php5.NewParser(src, v, withTokens)
	} else {
		parser = php7.NewParser(src, v, withTokens)
	}

	return parser, nil
}
