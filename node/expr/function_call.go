package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type FunctionCall struct {
	Function  node.Node
	Arguments []node.Node
}

func NewFunctionCall(Function node.Node, Arguments []node.Node) *FunctionCall {
	return &FunctionCall{
		Function,
		Arguments,
	}
}

func (n *FunctionCall) Attributes() map[string]interface{} {
	return nil
}

func (n *FunctionCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Function != nil {
		vv := v.GetChildrenVisitor("Function")
		n.Function.Walk(vv)
	}

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
