package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestStaticVar(t *testing.T) {
	src := `<? static $a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
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

func TestStaticVars(t *testing.T) {
	src := `<? static $a, $b = 1;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
						Expr:     &scalar.Lnumber{Value: "1"},
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

func TestStaticVars2(t *testing.T) {
	src := `<? static $a = 1, $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Static{
				Vars: []node.Node{
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
						Expr:     &scalar.Lnumber{Value: "1"},
					},
					&stmt.StaticVar{
						Variable: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
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
