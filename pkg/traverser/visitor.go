package traverser

import "github.com/z7zmey/php-parser/pkg/ast"

type Visitor interface {
	VisitNode(n ast.Node, depth int) bool
}
