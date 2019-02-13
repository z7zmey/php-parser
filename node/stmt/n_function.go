package stmt

import (
	"github.com/z7zmey/php-parser/freefloating"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Function node
type Function struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	ReturnsRef    bool
	PhpDocComment string
	FunctionName  node.Node
	Params        []node.Node
	ReturnType    node.Node
	Stmts         []node.Node
}

// NewFunction node constructor
func NewFunction(FunctionName node.Node, ReturnsRef bool, Params []node.Node, ReturnType node.Node, Stmts []node.Node, PhpDocComment string) *Function {
	return &Function{
		FreeFloating:  nil,
		ReturnsRef:    ReturnsRef,
		PhpDocComment: PhpDocComment,
		FunctionName:  FunctionName,
		Params:        Params,
		ReturnType:    ReturnType,
		Stmts:         Stmts,
	}
}

// SetPosition sets node position
func (n *Function) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Function) GetPosition() *position.Position {
	return n.Position
}

func (n *Function) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Function) Attributes() map[string]interface{} {
	// return n.attributes
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Function) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.FunctionName != nil {
		v.EnterChildNode("FunctionName", n)
		n.FunctionName.Walk(v)
		v.LeaveChildNode("FunctionName", n)
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
