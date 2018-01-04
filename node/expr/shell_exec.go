package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShellExec struct {
	attributes map[string]interface{}
	position   *node.Position
	parts      []node.Node
}

func NewShellExec(parts []node.Node) node.Node {
	return &ShellExec{
		map[string]interface{}{},
		nil,
		parts,
	}
}

func (n ShellExec) Attributes() map[string]interface{} {
	return n.attributes
}

func (n ShellExec) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n ShellExec) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n ShellExec) Position() *node.Position {
	return n.position
}

func (n ShellExec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n ShellExec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.parts != nil {
		vv := v.GetChildrenVisitor("parts")
		for _, nn := range n.parts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
