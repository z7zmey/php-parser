package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Constant struct {
	attributes   map[string]interface{}
	position     *node.Position
	ConstantName node.Node
	expr         node.Node
}

func NewConstant(ConstantName node.Node, expr node.Node, phpDocComment string) node.Node {
	return &Constant{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		ConstantName,
		expr,
	}
}

func (n Constant) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Constant) Position() *node.Position {
	return n.position
}

func (n Constant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Constant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		vv := v.GetChildrenVisitor("ConstantName")
		n.ConstantName.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
