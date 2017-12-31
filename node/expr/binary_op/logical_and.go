package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type LogicalAnd struct {
	BinaryOp
}

func NewLogicalAnd(variable node.Node, expression node.Node) node.Node {
	return LogicalAnd{
		BinaryOp{
			"BinaryLogicalAnd",
			map[string]interface{}{},
			nil,
			variable,
			expression,
		},
	}
}

func (n LogicalAnd) Name() string {
	return "LogicalAnd"
}

func (n LogicalAnd) Attributes() map[string]interface{} {
	return n.attributes
}

func (n LogicalAnd) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n LogicalAnd) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n LogicalAnd) Position() *node.Position {
	return n.position
}

func (n LogicalAnd) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n LogicalAnd) Walk(v node.Visitor) {
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
