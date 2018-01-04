package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Namespace struct {
	attributes    map[string]interface{}
	position      *node.Position
	namespaceName node.Node
	stmts         []node.Node
}

func NewNamespace(namespaceName node.Node, stmts []node.Node) node.Node {
	return Namespace{
		map[string]interface{}{},
		nil,
		namespaceName,
		stmts,
	}
}

func (n Namespace) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Namespace) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Namespace) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Namespace) Position() *node.Position {
	return n.position
}

func (n Namespace) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Namespace) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.namespaceName != nil {
		vv := v.GetChildrenVisitor("namespaceName")
		n.namespaceName.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
