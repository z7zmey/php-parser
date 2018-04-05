package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Do node
type Do struct {
	Stmt node.Node
	Cond node.Node
}

// NewDo node constructor
func NewDo(Stmt node.Node, Cond node.Node) *Do {
	return &Do{
		Stmt,
		Cond,
	}
}

// Attributes returns node attributes as map
func (n *Do) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Do) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	v.LeaveNode(n)
}
