package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Use struct {
	name    string
	useType node.Node
	use     node.Node
	alias   token.TokenInterface
}

func NewUse(useType node.Node, use node.Node, alias token.TokenInterface) node.Node {
	return Use{
		"Use",
		useType,
		use,
		alias,
	}
}

func (n Use) Name() string {
	return "Use"
}

func (n Use) SetType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n Use) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.Children("useType")
		n.useType.Walk(vv)
	}

	if n.use != nil {
		vv := v.Children("use")
		n.use.Walk(vv)
	}
}
