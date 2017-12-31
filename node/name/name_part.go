package name

import (
	"github.com/z7zmey/php-parser/node"
)

type NamePart struct {
	name       string
	attributes map[string]interface{}
}

func NewNamePart(value string) node.Node {
	return NamePart{
		"NamePart",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n NamePart) Name() string {
	return "NamePart"
}

func (n NamePart) Attributes() map[string]interface{} {
	return n.attributes
}

func (n NamePart) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n NamePart) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n NamePart) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
