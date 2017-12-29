package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Interface struct {
	name          string
	attributes    map[string]interface{}
	interfaceName node.Node
	extends       []node.Node
	stmts         []node.Node
}

func NewInterface(interfaceName node.Node, extends []node.Node, stmts []node.Node) node.Node {
	return Interface{
		"Interface",
		map[string]interface{}{},
		interfaceName,
		extends,
		stmts,
	}
}

func (n Interface) Name() string {
	return "Interface"
}

func (n Interface) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Interface) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.interfaceName != nil {
		vv := v.GetChildrenVisitor("interfaceName")
		n.interfaceName.Walk(vv)
	}

	if n.extends != nil {
		vv := v.GetChildrenVisitor("extends")
		for _, nn := range n.extends {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
