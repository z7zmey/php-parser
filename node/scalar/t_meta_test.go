package scalar_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&scalar.Dnumber{},
	&scalar.EncapsedStringPart{},
	&scalar.Encapsed{},
	&scalar.Heredoc{},
	&scalar.Lnumber{},
	&scalar.MagicConstant{},
	&scalar.String{},
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
