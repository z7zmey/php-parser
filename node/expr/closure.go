package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Closure struct {
	name       string
	attributes map[string]interface{}
	params     []node.Node
	uses       []node.Node
	returnType node.Node
	stmts      []node.Node
}

func NewClosure(params []node.Node, uses []node.Node, returnType node.Node, stmts []node.Node, isStatic bool, isReturnRef bool) node.Node {
	return Closure{
		"Closure",
		map[string]interface{}{
			"isReturnRef": isReturnRef,
			"isStatic":    isStatic,
		},
		params,
		uses,
		returnType,
		stmts,
	}
}

func (n Closure) Name() string {
	return "Closure"
}

func (n Closure) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Closure) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Closure) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Closure) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.params != nil {
		vv := v.GetChildrenVisitor("params")
		for _, nn := range n.params {
			nn.Walk(vv)
		}
	}

	if n.uses != nil {
		vv := v.GetChildrenVisitor("uses")
		for _, nn := range n.uses {
			nn.Walk(vv)
		}
	}

	if n.returnType != nil {
		vv := v.GetChildrenVisitor("returnType")
		n.returnType.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
