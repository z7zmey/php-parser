package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// AltIf node
type AltIf struct {
	Cond   node.Node
	Stmt   node.Node
	ElseIf []node.Node
	Else   node.Node
}

// NewAltIf node constructor
func NewAltIf(Cond node.Node, Stmt node.Node, ElseIf []node.Node, Else node.Node) *AltIf {
	return &AltIf{
		Cond,
		Stmt,
		ElseIf,
		Else,
	}
}

// Attributes returns node attributes as map
func (n *AltIf) Attributes() map[string]interface{} {
	return nil
}

// AddElseIf add AltElseIf node and returns AltIf node
func (n *AltIf) AddElseIf(ElseIf node.Node) node.Node {
	if n.ElseIf == nil {
		n.ElseIf = make([]node.Node, 0)
	}

	n.ElseIf = append(n.ElseIf, ElseIf)

	return n
}

// SetElse set AltElse node and returns AltIf node
func (n *AltIf) SetElse(Else node.Node) node.Node {
	n.Else = Else

	return n
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltIf) Walk(v walker.Visitor) {
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
