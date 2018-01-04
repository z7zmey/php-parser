package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StmtList struct {
	attributes map[string]interface{}
	position   *node.Position
	stmts      []node.Node
}

func NewStmtList(stmts []node.Node) node.Node {
	return &StmtList{
		map[string]interface{}{},
		nil,
		stmts,
	}
}

func (n StmtList) Attributes() map[string]interface{} {
	return n.attributes
}

func (n StmtList) Position() *node.Position {
	return n.position
}

func (n StmtList) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n StmtList) Walk(v node.Visitor) {
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
