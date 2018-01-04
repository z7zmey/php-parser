package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Namespace struct {
	attributes    map[string]interface{}
	position      *node.Position
	NamespaceName node.Node
	Stmts         []node.Node
}

func NewNamespace(NamespaceName node.Node, Stmts []node.Node) node.Node {
	return &Namespace{
		map[string]interface{}{},
		nil,
		NamespaceName,
		Stmts,
	}
}

func (n Namespace) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.NamespaceName != nil {
		vv := v.GetChildrenVisitor("NamespaceName")
		n.NamespaceName.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
