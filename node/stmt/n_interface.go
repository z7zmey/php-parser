package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Interface node
type Interface struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	PhpDocComment string
	InterfaceName node.Node
	Extends       *InterfaceExtends
	Stmts         []node.Node
}

// NewInterface node constructor
func NewInterface(InterfaceName node.Node, Extends *InterfaceExtends, Stmts []node.Node, PhpDocComment string) *Interface {
	return &Interface{
		FreeFloating:  nil,
		PhpDocComment: PhpDocComment,
		InterfaceName: InterfaceName,
		Extends:       Extends,
		Stmts:         Stmts,
	}
}

// SetPosition sets node position
func (n *Interface) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Interface) GetPosition() *position.Position {
	return n.Position
}

func (n *Interface) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Interface) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Interface) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.InterfaceName != nil {
		v.EnterChildNode("InterfaceName", n)
		n.InterfaceName.Walk(v)
		v.LeaveChildNode("InterfaceName", n)
	}

	if n.Extends != nil {
		v.EnterChildNode("Extends", n)
		n.Extends.Walk(v)
		v.LeaveChildNode("Extends", n)
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
