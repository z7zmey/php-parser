package binary_op

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type BinaryOp struct {
	position *node.Position
	comments []comment.Comment
	Left     node.Node
	Right    node.Node
}
