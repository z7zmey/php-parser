package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ElseIf struct {
	name string
	cond node.Node
	stmt node.Node
}

func NewElseIf(cond node.Node, stmt node.Node) node.Node {
	return ElseIf{
		"ElseIf",
		cond,
		stmt,
	}
}

func (n ElseIf) Name() string {
	return "ElseIf"
}

func (n ElseIf) Attributes() map[string]interface{} {
	return nil
}

func (n ElseIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
