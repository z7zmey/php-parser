package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type String struct {
	name       string
	attributes map[string]interface{}
}

func NewString(value string) node.Node {
	return String{
		"String",

		map[string]interface{}{
			"value": value,
		},
	}
}

func (n String) Name() string {
	return "String"
}

func (n String) Attributes() map[string]interface{} {
	return n.attributes
}

func (n String) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n String) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n String) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
