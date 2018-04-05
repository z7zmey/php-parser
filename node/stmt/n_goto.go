package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Goto node
type Goto struct {
	Label node.Node
}

// NewGoto node constructor
func NewGoto(Label node.Node) *Goto {
	return &Goto{
		Label,
	}
}

// Attributes returns node attributes as map
func (n *Goto) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Goto) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Label != nil {
		vv := v.GetChildrenVisitor("Label")
		n.Label.Walk(vv)
	}

	v.LeaveNode(n)
}
