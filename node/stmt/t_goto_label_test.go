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

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Label{
				LabelName: &node.Identifier{Value: "a"},
			},
			&stmt.Goto{
				Label: &node.Identifier{Value: "a"},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
