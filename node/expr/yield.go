package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type Yield struct {
	Key   node.Node
	Value node.Node
}

func NewYield(Key node.Node, Value node.Node) *Yield {
	return &Yield{
		Key,
		Value,
	}
}

func (n *Yield) Attributes() map[string]interface{} {
	return nil
}

func (n *Yield) Walk(v node.Visitor) {
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
