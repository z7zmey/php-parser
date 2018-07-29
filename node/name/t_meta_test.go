package name_test

import (
	"testing"

	"github.com/z7zmey/php-parser/node/name"

	"github.com/z7zmey/php-parser/meta"

	"github.com/z7zmey/php-parser/node"
)

var nodes = []node.Node{
	&name.FullyQualified{},
	&name.NamePart{},
	&name.Name{},
	&name.Relative{},
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
