package main

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/errors"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
)

type Parser interface {
	Parse() int
	GetPath() string
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	GetComments() comment.Comments
	GetPositions() position.Positions
}
