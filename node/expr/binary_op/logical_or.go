package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalOr struct {
	BinaryOp
}

func NewLogicalOr(variable node.Node, expression node.Node) node.Node {
	return LogicalOr{
		BinaryOp{
			"BinaryLogicalOr",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n LogicalOr) Name() string {
	return "LogicalOr"
}

func (n LogicalOr) Attributes() map[string]interface{} {
	return n.attributes
}

func (n LogicalOr) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n LogicalOr) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n LogicalOr) Position() *node.Position {
	return n.position
}

func (n LogicalOr) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n LogicalOr) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.left != nil {
		vv := v.GetChildrenVisitor("left")
		n.left.Walk(vv)
	}

	if n.right != nil {
		vv := v.GetChildrenVisitor("right")
		n.right.Walk(vv)
	}

	v.LeaveNode(n)
}
