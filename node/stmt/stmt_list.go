package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type StmtList struct {
	attributes map[string]interface{}
	position   *node.Position
	Stmts      []node.Node
}

func NewStmtList(Stmts []node.Node) node.Node {
	return StmtList{
		map[string]interface{}{},
		nil,
		Stmts,
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

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
