package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// AltElse node
type AltElse struct {
	Stmt node.Node
}

// NewAltElse node constructor
func NewAltElse(Stmt node.Node) *AltElse {
	return &AltElse{
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *AltElse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *AltElse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
