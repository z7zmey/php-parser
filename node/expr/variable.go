package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Variable struct {
	name       string
	attributes map[string]interface{}
	varName    node.Node
}

func NewVariable(varName node.Node) node.Node {
	return Variable{
		"Variable",
		map[string]interface{}{},
		varName,
	}
}

func (n Variable) Name() string {
	return "Variable"
}

func (n Variable) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.varName != nil {
		vv := v.GetChildrenVisitor("varName")
		n.varName.Walk(vv)
	}

	v.LeaveNode(n)
}
