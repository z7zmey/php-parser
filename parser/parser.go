package parser

import (
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
)

// Parser interface
type Parser interface {
	Parse() int
	GetPath() string
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	GetComments() Comments
	GetPositions() Positions
}
