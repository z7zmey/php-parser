package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type For struct {
	attributes map[string]interface{}
	position   *node.Position
	init       []node.Node
	cond       []node.Node
	loop       []node.Node
	stmt       node.Node
}

func NewFor(init []node.Node, cond []node.Node, loop []node.Node, stmt node.Node) node.Node {
	return &For{
		map[string]interface{}{},
		nil,
		init,
		cond,
		loop,
		stmt,
	}
}

func (n For) Attributes() map[string]interface{} {
	return n.attributes
}

func (n For) Position() *node.Position {
	return n.position
}

func (n For) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n For) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.init != nil {
		vv := v.GetChildrenVisitor("init")
		for _, nn := range n.init {
			nn.Walk(vv)
		}
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		for _, nn := range n.cond {
			nn.Walk(vv)
		}
	}

	if n.loop != nil {
		vv := v.GetChildrenVisitor("loop")
		for _, nn := range n.loop {
			nn.Walk(vv)
		}
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
