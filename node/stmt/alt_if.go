package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type AltIf struct {
	position *node.Position
	comments *[]comment.Comment
	Cond     node.Node
	Stmt     node.Node
	ElseIf   []node.Node
	_else    node.Node
}

func NewAltIf(Cond node.Node, Stmt node.Node) *AltIf {
	return &AltIf{
		nil,
		nil,
		Cond,
		Stmt,
		nil,
		nil,
	}
}

func (n *AltIf) Attributes() map[string]interface{} {
	return nil
}

func (n *AltIf) Position() *node.Position {
	return n.position
}

func (n *AltIf) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *AltIf) Comments() *[]comment.Comment {
	return n.comments
}

func (n *AltIf) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *AltIf) AddElseIf(ElseIf node.Node) node.Node {
	if n.ElseIf == nil {
		n.ElseIf = make([]node.Node, 0)
	}

	n.ElseIf = append(n.ElseIf, ElseIf)

	return n
}

func (n *AltIf) SetElse(_else node.Node) node.Node {
	n._else = _else

	return n
}

func (n *AltIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	if n.ElseIf != nil {
		vv := v.GetChildrenVisitor("ElseIf")
		for _, nn := range n.ElseIf {
			nn.Walk(vv)
		}
	}

	if n._else != nil {
		vv := v.GetChildrenVisitor("else")
		n._else.Walk(vv)
	}

	v.LeaveNode(n)
}
