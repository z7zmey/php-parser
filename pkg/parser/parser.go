package parser

import (
	"sync"

	"github.com/z7zmey/php-parser/internal/graph"
	"github.com/z7zmey/php-parser/internal/parser"
	"github.com/z7zmey/php-parser/internal/parser/php7"
	"github.com/z7zmey/php-parser/pkg/errors"
	"github.com/z7zmey/php-parser/pkg/traverser"
)

var traverserPool = sync.Pool{
	New: func() interface{} { return new(graph.Graph) },
}

type Parser interface {
	Parse([]byte) (traverser.Traverser, []*errors.Error)
	WithTokens() Parser
}

type phpParser struct {
	parser parser.Parser
}

func (p *phpParser) Parse(data []byte) (traverser.Traverser, []*errors.Error) {
	t := traverserPool.New().(*graph.Graph)
	p.parser.Parse(data, t)
	return t, p.parser.GetErrors()
}

func (p *phpParser) WithTokens() Parser {
	p.parser.WithTokens()
	return p
}

func NewPHP7Parser() Parser {
	return &phpParser{
		parser: php7.NewParser(),
	}
}

func Reuse(t traverser.Traverser) {
	traverserPool.Put(t)
}
