package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Catch struct {
	attributes map[string]interface{}
	position   *node.Position
	types      []node.Node
	variable   node.Node
	stmts      []node.Node
}

func NewCatch(types []node.Node, variable node.Node, stmts []node.Node) node.Node {
	return &Catch{
		map[string]interface{}{},
		nil,
		types,
		variable,
		stmts,
	}
}

func (n Catch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Catch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Catch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Catch) Position() *node.Position {
	return n.position
}

func (n Catch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Catch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.types != nil {
		vv := v.GetChildrenVisitor("types")
		for _, nn := range n.types {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
