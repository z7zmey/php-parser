package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestGotoLabel(t *testing.T) {
	src := `<? a: goto a;`

	expected := &node.Root{
		Stmts: []node.Node{
			&stmt.Label{
				LabelName: &node.Identifier{Value: "a"},
			},
			&stmt.Goto{
				Label: &node.Identifier{Value: "a"},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assertEqual(t, expected, actual)

	php5parser := php5.NewParser(bytes.NewBufferString(src), "test.php")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assertEqual(t, expected, actual)
}
