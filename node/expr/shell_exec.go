package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ShellExec) Name() string {
	return "ShellExec"
}

type ShellExec struct {
	name  string
	parts []node.Node
}

func NewShellExec(parts []node.Node) node.Node {
	return ShellExec{
		"ShellExec",
		parts,
	}
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
