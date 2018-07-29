package assign_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&assign.Reference{},
	&assign.Assign{},
	&assign.BitwiseAnd{},
	&assign.BitwiseOr{},
	&assign.BitwiseXor{},
	&assign.Concat{},
	&assign.Div{},
	&assign.Minus{},
	&assign.Mod{},
	&assign.Mul{},
	&assign.Plus{},
	&assign.Pow{},
	&assign.ShiftLeft{},
	&assign.ShiftRight{},
	&assign.ShiftRight{},
}

func TestMeta(t *testing.T) {
	expected := &meta.Collection{
		&meta.Data{
			Value:    "//comment\n",
			Type:     meta.CommentType,
			Position: nil,
		},
		&meta.Data{
			Value:    "    ",
			Type:     meta.WhiteSpaceType,
			Position: nil,
		},
	}
	for _, n := range nodes {
		n.GetMeta().Push(*expected...)
		actual := n.GetMeta()
		assertEqual(t, expected, actual)
	}
}
