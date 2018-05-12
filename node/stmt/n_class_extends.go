package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClassExtends node
type ClassExtends struct {
	ClassName node.Node
}

// NewClassExtends node constructor
func NewClassExtends(className node.Node) *ClassExtends {
	return &ClassExtends{
		className,
	}
}

// Attributes returns node attributes as map
func (n *ClassExtends) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassExtends) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ClassName != nil {
		vv := v.GetChildrenVisitor("ClassName")
		n.ClassName.Walk(vv)
	}

	v.LeaveNode(n)
}
