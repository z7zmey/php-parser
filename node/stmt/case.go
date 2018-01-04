package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Case struct {
	attributes map[string]interface{}
	position   *node.Position
	cond       node.Node
	stmts      []node.Node
}

func NewCase(cond node.Node, stmts []node.Node) node.Node {
	return &Case{
		map[string]interface{}{},
		nil,
		cond,
		stmts,
	}
}

func (n Case) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Case) Position() *node.Position {
	return n.position
}

func (n Case) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Case) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
