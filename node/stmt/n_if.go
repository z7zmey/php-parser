package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// If node
type If struct {
	Cond   node.Node
	Stmt   node.Node
	ElseIf []node.Node
	Else   node.Node
}

// NewIf node constructor
func NewIf(Cond node.Node, Stmt node.Node, ElseIf []node.Node, Else node.Node) *If {
	return &If{
		Cond,
		Stmt,
		ElseIf,
		Else,
	}
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
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Else != nil {
		vv := v.GetChildrenVisitor("Else")
		n.Else.Walk(vv)
	}

	v.LeaveNode(n)
}
