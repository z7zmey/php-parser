package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Do struct {
	name string
	stmt node.Node
	cond node.Node
}

func NewDo(stmt node.Node, cond node.Node) node.Node {
	return Do{
		"Do",
		stmt,
		cond,
	}
}

func (n Do) Name() string {
	return "Do"
}

func (n Do) Attributes() map[string]interface{} {
	return nil
}

func (n Do) Walk(v node.Visitor) {
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
