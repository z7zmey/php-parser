package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Default struct {
	name  string
	stmts []node.Node
}

func NewDefault(stmts []node.Node) node.Node {
	return Default{
		"Default",
		stmts,
	}
}

func (n Default) Name() string {
	return "Default"
}

func (n Default) Attributes() map[string]interface{} {
	return nil
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
