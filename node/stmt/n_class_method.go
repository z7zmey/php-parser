package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// ClassMethod node
type ClassMethod struct {
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
		ReturnsRef,
		PhpDocComment,
		MethodName,
		Modifiers,
		Params,
		ReturnType,
		Stmt,
	}
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
		vv := v.GetChildrenVisitor("MethodName")
		n.MethodName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Params != nil {
		vv := v.GetChildrenVisitor("Params")
		for _, nn := range n.Params {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.ReturnType != nil {
		vv := v.GetChildrenVisitor("ReturnType")
		n.ReturnType.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
