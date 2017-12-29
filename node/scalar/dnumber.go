package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Dnumber struct {
	name       string
	attributes map[string]interface{}
}

func NewDnumber(value string) node.Node {
	return Dnumber{
		"Dnumber",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n Dnumber) Name() string {
	return "Dnumber"
}

func (n Dnumber) Attributes() map[string]interface{} {
	return nil
}

func (n Dnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
