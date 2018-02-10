package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/scalar"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestArray(t *testing.T) {
	src := `<? array();`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestArrayItem(t *testing.T) {
	src := `<? array(1);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							ByRef: false,
							Val:   &scalar.Lnumber{Value: "1"},
						},
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

func TestArrayItems(t *testing.T) {
	src := `<? array(1=>1, &$b,);`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &expr.Array{
					Items: []node.Node{
						&expr.ArrayItem{
							ByRef: false,
							Key:   &scalar.Lnumber{Value: "1"},
							Val:   &scalar.Lnumber{Value: "1"},
						},
						&expr.ArrayItem{
							ByRef: true,
							Val:   &expr.Variable{VarName: &node.Identifier{Value: "$b"}},
						},
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
