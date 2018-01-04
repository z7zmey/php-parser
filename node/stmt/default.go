package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Default struct {
	attributes map[string]interface{}
	position   *node.Position
	stmts      []node.Node
}

func NewDefault(stmts []node.Node) node.Node {
	return &Default{
		map[string]interface{}{},
		nil,
		stmts,
	}
}

func (n Default) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Default) Position() *node.Position {
	return n.position
}

func (n Default) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Default) Walk(v node.Visitor) {
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
