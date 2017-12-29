package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type FunctionCall struct {
	name       string
	attributes map[string]interface{}
	function   node.Node
	arguments  []node.Node
}

func NewFunctionCall(function node.Node, arguments []node.Node) node.Node {
	return FunctionCall{
		"FunctionCall",
		map[string]interface{}{},
		function,
		arguments,
	}
}

func (n FunctionCall) Name() string {
	return "FunctionCall"
}

func (n FunctionCall) Attributes() map[string]interface{} {
	return n.attributes
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
