package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Closure struct {
	attributes map[string]interface{}
	position   *node.Position
	params     []node.Node
	uses       []node.Node
	returnType node.Node
	stmts      []node.Node
}

func NewClosure(params []node.Node, uses []node.Node, returnType node.Node, stmts []node.Node, isStatic bool, isReturnRef bool, phpDocComment string) node.Node {
	return Closure{
		map[string]interface{}{
			"isReturnRef":   isReturnRef,
			"isStatic":      isStatic,
			"phpDocComment": phpDocComment,
		},
		nil,
		params,
		uses,
		returnType,
		stmts,
	}
}

func (n Closure) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Closure) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Closure) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Closure) Position() *node.Position {
	return n.position
}

func (n Closure) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
