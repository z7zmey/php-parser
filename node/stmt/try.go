package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Try struct {
	attributes map[string]interface{}
	position   *node.Position
	stmts      []node.Node
	catches    []node.Node
	finally    node.Node
}

func NewTry(stmts []node.Node, catches []node.Node, finally node.Node) node.Node {
	return &Try{
		map[string]interface{}{},
		nil,
		stmts,
		catches,
		finally,
	}
}

func (n Try) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Try) Position() *node.Position {
	return n.position
}

func (n Try) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Try) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	if n.catches != nil {
		vv := v.GetChildrenVisitor("catches")
		for _, nn := range n.catches {
			nn.Walk(vv)
		}
	}

	if n.finally != nil {
		vv := v.GetChildrenVisitor("finally")
		n.finally.Walk(vv)
	}

	v.LeaveNode(n)
}
