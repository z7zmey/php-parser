package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type Catch struct {
	position *node.Position
	comments []comment.Comment
	Types    []node.Node
	Variable node.Node
	Stmts    []node.Node
}

func NewCatch(Types []node.Node, Variable node.Node, Stmts []node.Node) *Catch {
	return &Catch{
		nil,
		nil,
		Types,
		Variable,
		Stmts,
	}
}

func (n *Catch) Attributes() map[string]interface{} {
	return nil
}

func (n *Catch) Position() *node.Position {
	return n.position
}

func (n *Catch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *Catch) Comments() []comment.Comment {
	return n.comments
}

func (n *Catch) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *Catch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Types != nil {
		vv := v.GetChildrenVisitor("Types")
		for _, nn := range n.Types {
			nn.Walk(vv)
		}
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
