package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type ShellExec struct {
	position *node.Position
	Parts    []node.Node
}

func NewShellExec(Parts []node.Node) *ShellExec {
	return &ShellExec{
		nil,
		Parts,
	}
}

func (n *ShellExec) Attributes() map[string]interface{} {
	return nil
}

func (n *ShellExec) Position() *node.Position {
	return n.position
}

func (n *ShellExec) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *ShellExec) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
