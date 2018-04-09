package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestInlineHtml(t *testing.T) {
	src := `<? ?> <div></div>`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Nop{},
			&stmt.InlineHtml{Value: "<div></div>"},
		},
	}

	actual, _, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
