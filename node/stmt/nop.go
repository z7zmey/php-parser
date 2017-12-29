package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Nop struct {
	name       string
	attributes map[string]interface{}
}

func NewNop() node.Node {
	return Nop{
		"Nop",
		map[string]interface{}{},
	}
}

func (n Nop) Name() string {
	return "Nop"
}

func (n Nop) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
