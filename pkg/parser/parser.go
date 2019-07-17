package parser

import (
	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/parser"
	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

type Parser interface {
	Parse([]byte) (traverser.Traverser, []*errors.Error)
}

type phpParser struct {
	p parser.Parser
	t *graph.AST
}

func (p phpParser) Parse(data []byte) (traverser.Traverser, []*errors.Error) {
	p.p.Parse(data, p.t)
	return p.t, p.p.GetErrors()
}

func NewPHP7Parser() Parser {
	return phpParser{
		p: php7.NewParser(),
		t: new(graph.AST),
	}
}
