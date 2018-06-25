package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// If node
type If struct {
	Comments []*comment.Comment
	Position *position.Position
	Cond     node.Node
	Stmt     node.Node
	ElseIf   []node.Node
	Else     node.Node
}

// NewIf node constructor
func NewIf(Cond node.Node, Stmt node.Node, ElseIf []node.Node, Else node.Node) *If {
	return &If{
		Cond:   Cond,
		Stmt:   Stmt,
		ElseIf: ElseIf,
		Else:   Else,
	}
}

// SetPosition sets node position
func (n *If) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *If) GetPosition() *position.Position {
	return n.Position
}

func (n *If) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *If) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *If) Attributes() map[string]interface{} {
	return nil
}

// AddElseIf add ElseIf node and returns AltIf node
func (n *If) AddElseIf(ElseIf node.Node) node.Node {
	if n.ElseIf == nil {
		n.ElseIf = make([]node.Node, 0)
	}

	n.ElseIf = append(n.ElseIf, ElseIf)

	return n
}

// SetElse set Else node and returns AltIf node
func (n *If) SetElse(Else node.Node) node.Node {
	n.Else = Else

	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *If) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Cond != nil {
		v.EnterChildNode("Cond", n)
		n.Cond.Walk(v)
		v.LeaveChildNode("Cond", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	if n.ElseIf != nil {
		v.EnterChildList("ElseIf", n)
		for _, nn := range n.ElseIf {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("ElseIf", n)
	}

	if n.Else != nil {
		v.EnterChildNode("Else", n)
		n.Else.Walk(v)
		v.LeaveChildNode("Else", n)
	}

	v.LeaveNode(n)
}
