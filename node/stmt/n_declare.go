package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Declare node
type Declare struct {
	Consts []node.Node
	Stmt   node.Node
}

// NewDeclare node constructor
func NewDeclare(Consts []node.Node, Stmt node.Node) *Declare {
	return &Declare{
		Consts,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *Declare) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Declare) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Consts != nil {
		vv := v.GetChildrenVisitor("Consts")
		for _, nn := range n.Consts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
