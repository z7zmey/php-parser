package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	attributes map[string]interface{}
	position   *node.Position
	VarName    node.Node
}

func NewVariable(VarName node.Node) node.Node {
	return &Variable{
		map[string]interface{}{},
		nil,
		VarName,
	}
}

func (n Variable) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Variable) Position() *node.Position {
	return n.position
}

func (n Variable) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		vv := v.GetChildrenVisitor("VarName")
		n.VarName.Walk(vv)
	}

	v.LeaveNode(n)
}
