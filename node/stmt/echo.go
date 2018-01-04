package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Echo struct {
	attributes map[string]interface{}
	position   *node.Position
	exprs      []node.Node
}

func NewEcho(exprs []node.Node) node.Node {
	return &Echo{
		map[string]interface{}{},
		nil,
		exprs,
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

	if n.exprs != nil {
		vv := v.GetChildrenVisitor("exprs")
		for _, nn := range n.exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
