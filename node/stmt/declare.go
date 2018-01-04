package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Declare struct {
	attributes map[string]interface{}
	position   *node.Position
	consts     []node.Node
	stmt       node.Node
}

func NewDeclare(consts []node.Node, stmt node.Node) node.Node {
	return &Declare{
		map[string]interface{}{},
		nil,
		consts,
		stmt,
	}
}

func (n Declare) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Declare) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Declare) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Declare) Position() *node.Position {
	return n.position
}

func (n Declare) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Declare) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.consts != nil {
		vv := v.GetChildrenVisitor("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
