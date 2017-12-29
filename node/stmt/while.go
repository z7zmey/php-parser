package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type While struct {
	name       string
	attributes map[string]interface{}
	token      token.Token
	cond       node.Node
	stmt       node.Node
}

func NewWhile(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return While{
		"While",
		map[string]interface{}{},
		token,
		cond,
		stmt,
	}
}

func (n While) Name() string {
	return "While"
}

func (n While) Attributes() map[string]interface{} {
	return n.attributes
}

func (n While) Walk(v node.Visitor) {
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
