package stmt_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
)

func TestGlobal(t *testing.T) {
	src := `<? global $a;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Global{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "$a"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)

	actual, _, _ = php5.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestGlobalVars(t *testing.T) {
	src := `<? global $a, $b, $$c, ${foo()};`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Global{
				Vars: []node.Node{
					&expr.Variable{VarName: &node.Identifier{Value: "$a"}},
					&expr.Variable{VarName: &node.Identifier{Value: "$b"}},
					&expr.Variable{VarName: &expr.Variable{VarName: &node.Identifier{Value: "$c"}}},
					&expr.Variable{
						VarName: &expr.FunctionCall{
							Function: &name.Name{
								Parts: []node.Node{
									&name.NamePart{Value: "foo"},
								},
							},
							Arguments: []node.Node{},
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
