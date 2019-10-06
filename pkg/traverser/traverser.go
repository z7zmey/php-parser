package traverser

import "github.com/z7zmey/php-parser/pkg/ast"

type Traverser interface {
	Traverse(v Visitor)
	RootNode() *ast.Node
}
