package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Lnumber struct {
	Value string
}

func NewLnumber(Value string) *Lnumber {
	return &Lnumber{
		Value,
	}
}

func (n *Lnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
