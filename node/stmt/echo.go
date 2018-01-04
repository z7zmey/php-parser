package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Echo struct {
	attributes map[string]interface{}
	position   *node.Position
	Exprs      []node.Node
}

func NewEcho(Exprs []node.Node) node.Node {
	return &Echo{
		map[string]interface{}{},
		nil,
		Exprs,
	}
}

func (n Echo) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Echo) Position() *node.Position {
	return n.position
}

func (n Echo) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Echo) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		vv := v.GetChildrenVisitor("Exprs")
		for _, nn := range n.Exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
