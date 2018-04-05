package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Variable node
type Variable struct {
	VarName node.Node
}

// NewVariable node constructor
func NewVariable(VarName node.Node) *Variable {
	return &Variable{
		VarName,
	}
}

// Attributes returns node attributes as map
func (n *Variable) Attributes() map[string]interface{} {
	return nil
}

// SetVarName reset var name
func (n *Variable) SetVarName(VarName node.Node) {
	n.VarName = VarName
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Variable) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VarName != nil {
		vv := v.GetChildrenVisitor("VarName")
		n.VarName.Walk(vv)
	}

	v.LeaveNode(n)
}
