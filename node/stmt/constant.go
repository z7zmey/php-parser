package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Constant struct {
	position      *node.Position
	PhpDocComment string
	ConstantName  node.Node
	Expr          node.Node
}

func NewConstant(ConstantName node.Node, Expr node.Node, PhpDocComment string) *Constant {
	return &Constant{
		nil,
		PhpDocComment,
		ConstantName,
		Expr,
	}
}

func (n *Constant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n *Constant) Position() *node.Position {
	return n.position
}

func (n *Constant) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Constant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		vv := v.GetChildrenVisitor("ConstantName")
		n.ConstantName.Walk(vv)
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
