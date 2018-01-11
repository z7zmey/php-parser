package expr

import (
	"github.com/z7zmey/php-parser/node"
)

// Variable node
type Variable struct {
	VarName node.Node
}

// NewVariable node constuctor
func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		VarName,
	}
}

// Attributes returns node attributes as map
func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Variable) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		vv := v.GetChildrenVisitor("VarName")
		n.VarName.Walk(vv)
	}

	v.LeaveNode(n)
}
