package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n MethodCall) Name() string {
	return "MethodCall"
}

type MethodCall struct {
	name      string
	variable  node.Node
	method    node.Node
	arguments []node.Node
}

func NewMethodCall(variable node.Node, method node.Node, arguments []node.Node) node.Node {
	return MethodCall{
		"MethodCall",
		variable,
		method,
		arguments,
	}
}

func (n MethodCall) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.Children("variable")
		n.variable.Walk(vv)
	}

	if n.method != nil {
		vv := v.Children("method")
		n.method.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.Children("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}
}
