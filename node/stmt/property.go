package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Property struct {
	attributes map[string]interface{}
	position   *node.Position
	Variable   node.Node
	expr       node.Node
}

func NewProperty(Variable node.Node, expr node.Node, phpDocComment string) node.Node {
	return &Property{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		Variable,
		expr,
	}
}
func (n Property) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Property) Position() *node.Position {
	return n.position
}

func (n Property) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Property) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
