package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ClassExtends node
type ClassExtends struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	ClassName    node.Node
}

// NewClassExtends node constructor
func NewClassExtends(className node.Node) *ClassExtends {
	return &ClassExtends{
		FreeFloating: nil,
		ClassName:    className,
	}
}

// SetPosition sets node position
func (n *ClassExtends) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClassExtends) GetPosition() *position.Position {
	return n.Position
}

func (n *ClassExtends) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
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
		v.EnterChildNode("ClassName", n)
		n.ClassName.Walk(v)
		v.LeaveChildNode("ClassName", n)
	}

	v.LeaveNode(n)
}
