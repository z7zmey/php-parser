package expr

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// ArrowFunction node
type ArrowFunction struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	ReturnsRef    bool
	Static        bool
	PhpDocComment string
	Params        []node.Node
	ReturnType    node.Node
	Expr          node.Node
}

// NewArrowFunction node constructor
func NewArrowFunction(Params []node.Node, ReturnType node.Node, Stmt node.Node, Static bool, ReturnsRef bool, PhpDocComment string) *ArrowFunction {
	return &ArrowFunction{
		FreeFloating:  nil,
		ReturnsRef:    ReturnsRef,
		Static:        Static,
		PhpDocComment: PhpDocComment,
		Params:        Params,
		ReturnType:    ReturnType,
		Expr:          Stmt,
	}
}

// SetPosition sets node position
func (n *ArrowFunction) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ArrowFunction) GetPosition() *position.Position {
	return n.Position
}

func (n *ArrowFunction) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ArrowFunction) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"Static":        n.Static,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ArrowFunction) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
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

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
