package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Else node
type Else struct {
	Stmt node.Node
}

// NewElse node constructor
func NewElse(Stmt node.Node) *Else {
	return &Else{
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *Else) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Else) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
