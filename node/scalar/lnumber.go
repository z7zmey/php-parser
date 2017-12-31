package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Lnumber struct {
	name       string
	attributes map[string]interface{}
}

func NewLnumber(value string) node.Node {
	return Lnumber{
		"Lnumber",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n Lnumber) Name() string {
	return "Lnumber"
}

func (n Lnumber) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Lnumber) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Lnumber) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
