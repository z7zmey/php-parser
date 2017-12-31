package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Try struct {
	name       string
	attributes map[string]interface{}
	stmts      []node.Node
	catches    []node.Node
	finally    node.Node
}

func NewTry(stmts []node.Node, catches []node.Node, finally node.Node) node.Node {
	return Try{
		"Try",
		map[string]interface{}{},
		stmts,
		catches,
		finally,
	}
}

func (n Try) Name() string {
	return "Try"
}

func (n Try) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Try) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Try) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Try) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	if n.catches != nil {
		vv := v.GetChildrenVisitor("catches")
		for _, nn := range n.catches {
			nn.Walk(vv)
		}
	}

	if n.finally != nil {
		vv := v.GetChildrenVisitor("finally")
		n.finally.Walk(vv)
	}

	v.LeaveNode(n)
}
