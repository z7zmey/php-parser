package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n HaltCompiler) Name() string {
	return "HaltCompiler"
}

type HaltCompiler struct {
	name  string
	token token.Token
}

func NewHaltCompiler(token token.Token) node.Node {
	return HaltCompiler{
		"HaltCompiler",
		token,
	}
}

func (n HaltCompiler) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
