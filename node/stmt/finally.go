package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Finally struct {
	attributes map[string]interface{}
	position   *node.Position
	stmts      []node.Node
}

func NewFinally(stmts []node.Node) node.Node {
	return &Finally{
		map[string]interface{}{},
		nil,
		stmts,
	}
}

func (n Finally) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Finally) Position() *node.Position {
	return n.position
}

func (n Finally) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Finally) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
