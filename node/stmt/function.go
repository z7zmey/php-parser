package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Function struct {
	name         string
	attributes   map[string]interface{}
	position     *node.Position
	functionName node.Node
	params       []node.Node
	returnType   node.Node
	stmts        []node.Node
}

func NewFunction(functionName node.Node, returnsRef bool, params []node.Node, returnType node.Node, stmts []node.Node, phpDocComment string) node.Node {
	return Function{
		"Function",
		map[string]interface{}{
			"returnsRef":    returnsRef,
			"phpDocComment": phpDocComment,
		},
		nil,
		functionName,
		params,
		returnType,
		stmts,
	}
}

func (n Function) Name() string {
	return "Function"
}

func (n Function) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Function) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Function) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Function) Position() *node.Position {
	return n.position
}

func (n Function) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Function) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.functionName != nil {
		vv := v.GetChildrenVisitor("functionName")
		n.functionName.Walk(vv)
	}

	if n.params != nil {
		vv := v.GetChildrenVisitor("params")
		for _, nn := range n.params {
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
