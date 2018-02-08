package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestSimpleEcho(t *testing.T) {
	src := `<? echo $a, 1;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "$a"},
					},
					&scalar.Lnumber{Value: "1"},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestEcho(t *testing.T) {
	src := `<? echo($a);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Echo{
				Exprs: []node.Node{
					&expr.Variable{
						VarName: &node.Identifier{Value: "$a"},
					},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
