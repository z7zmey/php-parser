package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ClassMethod node
type ClassMethod struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	ReturnsRef    bool
	PhpDocComment string
	MethodName    node.Node
	Modifiers     []node.Node
	Params        []node.Node
	ReturnType    node.Node
	Stmt          node.Node
}

// NewClassMethod node constructor
func NewClassMethod(MethodName node.Node, Modifiers []node.Node, ReturnsRef bool, Params []node.Node, ReturnType node.Node, Stmt node.Node, PhpDocComment string) *ClassMethod {
	return &ClassMethod{
		FreeFloating:  nil,
		ReturnsRef:    ReturnsRef,
		PhpDocComment: PhpDocComment,
		MethodName:    MethodName,
		Modifiers:     Modifiers,
		Params:        Params,
		ReturnType:    ReturnType,
		Stmt:          Stmt,
	}
}

// SetPosition sets node position
func (n *ClassMethod) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClassMethod) GetPosition() *position.Position {
	return n.Position
}

func (n *ClassMethod) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ClassMethod) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClassMethod) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.MethodName != nil {
		v.EnterChildNode("MethodName", n)
		n.MethodName.Walk(v)
		v.LeaveChildNode("MethodName", n)
	}

	if n.Modifiers != nil {
		v.EnterChildList("Modifiers", n)
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Modifiers", n)
	}

	if n.Params != nil {
		v.EnterChildList("Params", n)
		for _, nn := range n.Params {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Params", n)
	}

	if n.ReturnType != nil {
		v.EnterChildNode("ReturnType", n)
		n.ReturnType.Walk(v)
		v.LeaveChildNode("ReturnType", n)
	}

	if n.Stmt != nil {
		v.EnterChildNode("Stmt", n)
		n.Stmt.Walk(v)
		v.LeaveChildNode("Stmt", n)
	}

	v.LeaveNode(n)
}
