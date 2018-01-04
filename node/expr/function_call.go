package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type FunctionCall struct {
	attributes map[string]interface{}
	position   *node.Position
	function   node.Node
	arguments  []node.Node
}

func NewFunctionCall(function node.Node, arguments []node.Node) node.Node {
	return &FunctionCall{
		map[string]interface{}{},
		nil,
		function,
		arguments,
	}
}

func (n FunctionCall) Attributes() map[string]interface{} {
	return n.attributes
}

func (n FunctionCall) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n FunctionCall) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n FunctionCall) Position() *node.Position {
	return n.position
}

func (n FunctionCall) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n FunctionCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.function != nil {
		vv := v.GetChildrenVisitor("function")
		n.function.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.GetChildrenVisitor("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
