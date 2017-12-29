package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Nop struct {
	name string
}

func NewNop() node.Node {
	return Nop{
		"Nop",
	}
}

func (n Nop) Name() string {
	return "Nop"
}

func (n Nop) Attributes() map[string]interface{} {
	return nil
}

func (n Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
