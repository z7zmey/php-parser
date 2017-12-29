package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type MagicConstant struct {
	name       string
	attributes map[string]interface{}
}

func NewMagicConstant(value string) node.Node {
	return MagicConstant{
		"MagicConstant",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n MagicConstant) Name() string {
	return "MagicConstant"
}

func (n MagicConstant) Attributes() map[string]interface{} {
	return n.attributes
}

func (n MagicConstant) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
