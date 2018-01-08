package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type MethodCall struct {
	position  *node.Position
	Variable  node.Node
	Method    node.Node
	Arguments []node.Node
}

func NewMethodCall(Variable node.Node, Method node.Node, Arguments []node.Node) *MethodCall {
	return &MethodCall{
		nil,
		Variable,
		Method,
		Arguments,
	}
}

func (n *MethodCall) Attributes() map[string]interface{} {
	return nil
}

func (n *MethodCall) Position() *node.Position {
	return n.position
}

func (n *MethodCall) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *MethodCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Method != nil {
		vv := v.GetChildrenVisitor("Method")
		n.Method.Walk(vv)
	}

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
