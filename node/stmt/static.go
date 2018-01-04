package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Static struct {
	attributes map[string]interface{}
	position   *node.Position
	vars       []node.Node
}

func NewStatic(vars []node.Node) node.Node {
	return &Static{
		map[string]interface{}{},
		nil,
		vars,
	}
}

func (n Static) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Static) Position() *node.Position {
	return n.position
}

func (n Static) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Static) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.GetChildrenVisitor("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
