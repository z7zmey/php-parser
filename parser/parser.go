package parser

import (
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/syntaxtree/linkedtree"
)

// Parser interface
type Parser interface {
	Parse([]byte, *linkedtree.AST) int
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	WithFreeFloating()
}
