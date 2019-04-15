package parser

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
)

// Parser interface
type Parser interface {
	Parse([]byte, *ast.AST) int
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	WithFreeFloating()
}
