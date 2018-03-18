package expr_test

import (
	"bytes"
	"testing"

	"github.com/z7zmey/php-parser/node/expr/assign"

	"github.com/z7zmey/php-parser/node/expr"

	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/stmt"
	"github.com/z7zmey/php-parser/php7"
)

func TestShortList(t *testing.T) {
	src := `<? [$a] = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestShortListArrayIndex(t *testing.T) {
	src := `<? [$a[]] = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.ArrayDimFetch{
									Variable: &expr.Variable{VarName: &node.Identifier{Value: "a"}},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}

func TestShortListList(t *testing.T) {
	src := `<? [list($a)] = $b;`

	expected := &stmt.StmtList{
		Stmts: []node.Node{
			&stmt.Expression{
				Expr: &assign.Assign{
					Variable: &expr.ShortList{
						Items: []node.Node{
							&expr.ArrayItem{
								ByRef: false,
								Val: &expr.List{
									Items: []node.Node{
										&expr.ArrayItem{
											ByRef: false,
											Val:   &expr.Variable{VarName: &node.Identifier{Value: "a"}},
										},
									},
								},
							},
						},
					},
					Expression: &expr.Variable{VarName: &node.Identifier{Value: "b"}},
				},
			},
		},
	}

	actual, _, _ := php7.Parse(bytes.NewBufferString(src), "test.php")
	assertEqual(t, expected, actual)
}
