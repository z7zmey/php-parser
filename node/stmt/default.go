package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Default struct {
	name       string
	attributes map[string]interface{}
	stmts      []node.Node
}

func NewDefault(stmts []node.Node) node.Node {
	return Default{
		"Default",
		map[string]interface{}{},
		stmts,
	}
}

func (n Default) Name() string {
	return "Default"
}

func (n Default) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Default) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Default) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Default) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
