package expr

import (
	"github.com/z7zmey/php-parser/node"
)

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

func (n MethodCall) Name() string {
	return "MethodCall"
}

func (n MethodCall) Attributes() map[string]interface{} {
	return nil
}

func (n MethodCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.method != nil {
		vv := v.GetChildrenVisitor("method")
		n.method.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.GetChildrenVisitor("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
