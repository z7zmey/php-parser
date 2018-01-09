package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Dnumber struct {
	Value string
}

func NewDnumber(Value string) *Dnumber {
	return &Dnumber{
		Value,
	}
}

func (n *Dnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *Dnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
