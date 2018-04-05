package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// AltElseIf node
type AltElseIf struct {
	Cond node.Node
	Stmt node.Node
}

// NewAltElseIf node constructor
func NewAltElseIf(Cond node.Node, Stmt node.Node) *AltElseIf {
	return &AltElseIf{
		Cond,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *AltElseIf) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltElseIf) Walk(v walker.Visitor) {
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
