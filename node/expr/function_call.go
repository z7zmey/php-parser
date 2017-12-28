package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n FunctionCall) Name() string {
	return "FunctionCall"
}

type FunctionCall struct {
	name      string
	function  node.Node
	arguments []node.Node
}

func NewFunctionCall(function node.Node, arguments []node.Node) node.Node {
	return FunctionCall{
		"FunctionCall",
		function,
		arguments,
	}
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
