package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestErrorSuppress(t *testing.T) {
	src := `<? @$a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.ErrorSuppress{
					Expr: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
