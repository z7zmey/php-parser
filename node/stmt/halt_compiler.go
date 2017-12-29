package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type HaltCompiler struct {
	name       string
	attributes map[string]interface{}
}

func NewHaltCompiler() node.Node {
	return HaltCompiler{
		"HaltCompiler",
		map[string]interface{}{},
	}
}

func (n HaltCompiler) Name() string {
	return "HaltCompiler"
}

func (n HaltCompiler) Attributes() map[string]interface{} {
	return n.attributes
}

func (n HaltCompiler) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
