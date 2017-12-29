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
	return nil
}

func (n Lnumber) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
