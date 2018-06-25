package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// UseList node
type UseList struct {
	Comments []*comment.Comment
	Position *position.Position
	UseType  node.Node
	Uses     []node.Node
}

// NewUseList node constructor
func NewUseList(UseType node.Node, Uses []node.Node) *UseList {
	return &UseList{
		UseType: UseType,
		Uses:    Uses,
	}
}

// SetPosition sets node position
func (n *UseList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *UseList) GetPosition() *position.Position {
	return n.Position
}

func (n *UseList) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *UseList) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *UseList) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UseList) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		v.EnterChildNode("UseType", n)
		n.UseType.Walk(v)
		v.LeaveChildNode("UseType", n)
	}

	if n.Uses != nil {
		v.EnterChildList("Uses", n)
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Uses", n)
	}

	v.LeaveNode(n)
}
