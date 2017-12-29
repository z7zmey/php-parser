package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type If struct {
	name       string
	attributes map[string]interface{}
	cond       node.Node
	stmt       node.Node
	elseIf     []node.Node
	_else      node.Node
}

func NewIf(cond node.Node, stmt node.Node) node.Node {
	return If{
		"If",
		map[string]interface{}{},
		cond,
		stmt,
		nil,
		nil,
	}
}

func (n If) Name() string {
	return "If"
}

func (n If) Attributes() map[string]interface{} {
	return n.attributes
}

func (n If) AddElseIf(elseIf node.Node) node.Node {
	if n.elseIf == nil {
		n.elseIf = make([]node.Node, 0)
	}

	n.elseIf = append(n.elseIf, elseIf)

	return n
}

func (n If) SetElse(_else node.Node) node.Node {
	n._else = _else

	return n
}

func (n If) Walk(v node.Visitor) {
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

	if n.elseIf != nil {
		vv := v.GetChildrenVisitor("elseIf")
		for _, nn := range n.elseIf {
			nn.Walk(vv)
		}
	}

	if n._else != nil {
		vv := v.GetChildrenVisitor("else")
		n._else.Walk(vv)
	}

	v.LeaveNode(n)
}
