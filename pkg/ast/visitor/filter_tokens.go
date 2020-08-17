package visitor

import (
	"github.com/z7zmey/php-parser/pkg/ast"
)

type FilterTokens struct {
	Null
}

func (v *FilterTokens) EnterNode(n ast.Vertex) bool {
	n.GetNode().Tokens = nil
	return true
}
