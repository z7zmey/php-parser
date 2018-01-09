package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type HaltCompiler struct {
}

func NewHaltCompiler() *HaltCompiler {
	return &HaltCompiler{}
}

func (n *HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

func (n *HaltCompiler) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
