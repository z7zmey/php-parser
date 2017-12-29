package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type PostInc struct {
	name     string
	variable node.Node
}

func NewPostInc(variable node.Node) node.Node {
	return PostInc{
		"PostInc",
		variable,
	}
}

func (n PostInc) Name() string {
	return "PostInc"
}

func (n PostInc) Attributes() map[string]interface{} {
	return nil
}

func (n PostInc) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	v.LeaveNode(n)
}
