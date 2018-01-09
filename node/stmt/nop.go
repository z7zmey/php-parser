package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Nop struct {
}

func NewNop() *Nop {
	return &Nop{}
}

func (n *Nop) Attributes() map[string]interface{} {
	return nil
}

func (n *Nop) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
