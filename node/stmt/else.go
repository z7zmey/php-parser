package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Else struct {
	name       string
	attributes map[string]interface{}
	stmt       node.Node
}

func NewElse(stmt node.Node) node.Node {
	return Else{
		"Else",
		map[string]interface{}{},
		stmt,
	}
}

func (n Else) Name() string {
	return "Else"
}

func (n Else) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Else) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
