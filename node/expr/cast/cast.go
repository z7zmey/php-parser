package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	position *node.Position
	Expr     node.Node
}
