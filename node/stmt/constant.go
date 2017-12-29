package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Constant struct {
	name         string
	constantName node.Node
	expr         node.Node
}

func NewConstant(constantName node.Node, expr node.Node) node.Node {
	return Constant{
		"Constant",
		constantName,
		expr,
	}
}

func (n Constant) Name() string {
	return "Constant"
}

func (n Constant) Attributes() map[string]interface{} {
	return nil
}

func (n Constant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.constantName != nil {
		vv := v.GetChildrenVisitor("constantName")
		n.constantName.Walk(vv)
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
