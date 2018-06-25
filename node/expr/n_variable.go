package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Variable node
type Variable struct {
	Comments []*comment.Comment
	Position *position.Position
	VarName  node.Node
}

// NewVariable node constructor
func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		VarName: VarName,
	}
}

// SetPosition sets node position
func (n *Variable) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Variable) GetPosition() *position.Position {
	return n.Position
}

func (n *Variable) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Variable) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

// SetVarName reset var name
func (n *Variable) SetVarName(VarName node.Node) {
	n.VarName = VarName
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Variable) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		v.EnterChildNode("VarName", n)
		n.VarName.Walk(v)
		v.LeaveChildNode("VarName", n)
	}

	v.LeaveNode(n)
}
