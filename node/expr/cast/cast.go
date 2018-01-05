package cast

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	position *node.Position
	comments *[]comment.Comment
	Expr     node.Node
}
