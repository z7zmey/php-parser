package traverser

import "github.com/z7zmey/php-parser/pkg/ast"

type Visitor interface {
	VisitNode(n ast.SimpleNode, depth int) bool
	VisitPosition(p ast.Position, depth int) bool
	VisitToken(t ast.Token, depth int) bool
}
