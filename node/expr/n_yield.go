package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Yield node
type Yield struct {
	Key   node.Node
	Value node.Node
}

// NewYield node constructor
func NewYield(Key node.Node, Value node.Node) *Yield {
	return &Yield{
		Key,
		Value,
	}
}

// Attributes returns node attributes as map
func (n *Yield) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Yield) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Value != nil {
		vv := v.GetChildrenVisitor("Value")
		n.Value.Walk(vv)
	}

	v.LeaveNode(n)
}
