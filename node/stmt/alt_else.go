package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type AltElse struct {
	name       string
	attributes map[string]interface{}
	stmt       node.Node
}

func NewAltElse(stmt node.Node) node.Node {
	return AltElse{
		"AltElse",
		map[string]interface{}{},
		stmt,
	}
}

func (n AltElse) Name() string {
	return "AltElse"
}

func (n AltElse) Attributes() map[string]interface{} {
	return n.attributes
}

func (n AltElse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
