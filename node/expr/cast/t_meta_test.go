package cast_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/expr/cast"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&cast.Array{},
	&cast.Bool{},
	&cast.Double{},
	&cast.Int{},
	&cast.Object{},
	&cast.String{},
	&cast.Unset{},
}

func TestMeta(t *testing.T) {
	expected := []meta.Meta{
		meta.NewComment("//comment\n", nil),
		meta.NewWhiteSpace("    ", nil),
	}
	for _, n := range nodes {
		n.AddMeta(expected)
		actual := n.GetMeta()
		assertEqual(t, expected, actual)
	}
}
