package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Label node
type Label struct {
	LabelName node.Node
}

// NewLabel node constructor
func NewLabel(LabelName node.Node) *Label {
	return &Label{
		LabelName,
	}
}

// Attributes returns node attributes as map
func (n *Label) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Label) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.LabelName != nil {
		vv := v.GetChildrenVisitor("LabelName")
		n.LabelName.Walk(vv)
	}

	v.LeaveNode(n)
}
