package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ElseIf node
type ElseIf struct {
	Cond node.Node
	Stmt node.Node
}

// NewElseIf node constructor
func NewElseIf(Cond node.Node, Stmt node.Node) *ElseIf {
	return &ElseIf{
		Cond,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *ElseIf) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ElseIf) Walk(v walker.Visitor) {
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

	v.LeaveNode(n)
}
